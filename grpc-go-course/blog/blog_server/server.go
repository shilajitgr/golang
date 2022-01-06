package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"reflect"

	"github.com/grpc-go-course/blog/blogpb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

var collection *mongo.Collection

type blogItem struct {
	ID       primitive.ObjectID `bson:"_id, omitempty`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

func (*server) CreateBlog(ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {

	blog := req.GetBlog()

	data := blogItem{
		AuthorID: blog.GetAuthorId(),
		Content:  blog.GetContent(),
		Title:    blog.GetTitle(),
	}
	fmt.Println("data to be inserted: \n", data)
	res, err := collection.InsertOne(ctx, data)
	fmt.Println("data inserted!")
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error occcured while inserting data: %v", err),
		)
	}

	objID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert result id to objectID"),
		)
	}

	return &blogpb.CreateBlogResponse{
		Blog: &blogpb.Blog{
			Id:       objID.Hex(),
			AuthorId: blog.GetAuthorId(),
			Content:  blog.GetContent(),
			Title:    blog.GetTitle(),
		},
	}, nil
}

func (*server) ReadBlog(ctx context.Context, req *blogpb.ReadBlogRequest) (*blogpb.ReadBlogResponse, error) {

	blogId := req.GetBlogId()

	objID, err := primitive.ObjectIDFromHex(fmt.Sprintf("%v", blogId))

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot Parse ID: %v", err),
		)
	}

	data := &blogItem{} // this empty struct will be filled with the match db entry data
	filter := bson.M{"_id": objID}

	findErr := collection.FindOne(context.Background(), filter).Decode(data)

	if findErr != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Could not Find/Decode matching data: %v", err),
		)
	}

	return &blogpb.ReadBlogResponse{
		Blog: &blogpb.Blog{
			Id:       data.ID.Hex(),
			AuthorId: data.AuthorID,
			Content:  data.Content,
			Title:    data.Title,
		},
	}, nil
}

func (*server) UpdateBlog(ctx context.Context, req *blogpb.UpdateBlogRequest) (*blogpb.UpdateBlogResponse, error) {

	blog := req.GetBlog()
	objId, err := primitive.ObjectIDFromHex(fmt.Sprintf("%v", blog.GetId()))
	fmt.Println("Data to update: \n", blog)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot Parse ID: %v", err),
		)
	}

	data := &blogItem{} // this empty struct will be filled with the match db entry data
	filter := bson.M{"_id": objId}

	findErr := collection.FindOne(context.Background(), filter).Decode(data)

	if findErr != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Could not Find/Decode matching data: %v", err),
		)
	}

	//updating data struct
	data.AuthorID = blog.GetAuthorId()
	data.Content = blog.GetContent()
	data.Title = blog.GetTitle()

	_, updateErr := collection.ReplaceOne(context.Background(), filter, data)

	if updateErr != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Failed to Update Blog: %v", err),
		)
	}

	return &blogpb.UpdateBlogResponse{
		Blog: &blogpb.Blog{
			Id:       data.ID.Hex(),
			AuthorId: data.AuthorID,
			Content:  data.Content,
			Title:    data.Title,
		},
	}, nil

}

func (*server) DeleteBlog(ctx context.Context, req *blogpb.DeleteBlogRequest) (*blogpb.DeleteBlogResponse, error) {

	blogId := req.GetBlogId()

	objID, err := primitive.ObjectIDFromHex(fmt.Sprintf("%v", blogId))

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot Parse ID: %v", err),
		)
	}

	filter := bson.M{"_id": objID}

	data := &blogItem{}

	findErr := collection.FindOne(context.Background(), filter).Decode(data)

	if findErr != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Could not Find/Decode matching data: %v", err),
		)
	}

	deleteErr := collection.FindOneAndDelete(context.Background(), filter) // DeleteOne() can also be used

	if deleteErr.Err() != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error while deleting entry: %v", deleteErr.Err()),
		)
	}

	return &blogpb.DeleteBlogResponse{
		BlogId: blogId,
	}, nil
}

func (*server) ListBlog(req *blogpb.ListBlogRequest, stream blogpb.BlogService_ListBlogServer) error {

	fmt.Println("ListBlog has been invoked.")
	cursor, err := collection.Find(context.Background(), primitive.D{{}})

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Server error: %v", err),
		)
	}

	defer cursor.Close(context.Background())
	// curson is pointer to db entries which needs to be closed at the end of the program

	for cursor.Next(context.Background()) {
		data := &blogItem{}
		decodeErr := cursor.Decode(data)
		if decodeErr != nil {
			return status.Errorf(
				codes.NotFound,
				fmt.Sprintf("Could not Decode data: %v", err),
			)
		}
		stream.Send(&blogpb.ListBlogResponse{
			Blog: &blogpb.Blog{
				Id:       data.ID.String(),
				AuthorId: data.AuthorID,
				Content:  data.Content,
				Title:    data.Title,
			},
		})

	}

	if cursor.Err() != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Server error: %v", cursor.Err()),
		)
	}
	return nil
}

func main() {
	log.SetFlags(log.Lshortfile | log.Lshortfile)

	// connecting to mongoDB
	fmt.Println("Connecting to mongoDB...")

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("goDB").Collection("Blog") // assigning value to the global var "collection"

	fmt.Printf("Type of collection object: %v", reflect.TypeOf(collection))

	fmt.Println("Starting Blog service")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Server could not listen %v", err)
	}

	s := grpc.NewServer()
	blogpb.RegisterBlogServiceServer(s, &server{})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server could not start %v", err)
		}
	}()

	// wait for ctrl+c to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt) // os.Interrupt refers to ctrl+c

	// Keep server running until ctrl+c is pressed, i.e., os.Interrupt signal is received
	<-ch

	fmt.Println("Stopping the server...")
	s.Stop()
	fmt.Println("Closing the listener...")
	lis.Close()
	fmt.Println("Closing mongoDB connection")
	client.Disconnect(context.TODO())
	fmt.Println("End of Program")
}
