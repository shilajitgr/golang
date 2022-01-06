package main

import (
	"context"
	"fmt"
	"log"

	"github.com/grpc-go-course/blog/blogpb"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}
	defer conn.Close()
	c := blogpb.NewBlogServiceClient(conn)

	// blog := &blogpb.Blog{
	// 	AuthorId: "Shilajit Guha Roy",
	// 	Content:  "Content of my first blog",
	// 	Title:    "My first blog",
	// }

	// res, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Blog has been created: %v", res)

	// // readying blog

	// fmt.Println("Reading blog")
	// _, err = c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "314wsydj"})
	// if err != nil {
	// 	fmt.Println("Error while fetching details.")
	// }

	// readRes, readErr := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: res.GetBlog().GetId()})

	// if readErr != nil {
	// 	fmt.Println("Error while fetching details", readErr)
	// }

	// fmt.Printf("Response received: %v", readRes)

	// //updating a new blog

	// blog2 := &blogpb.Blog{
	// 	Id:       res.GetBlog().GetId(),
	// 	AuthorId: "Google",
	// 	Content:  "Content of my second company",
	// 	Title:    "Their first blog",
	// }

	// updateRes, updaterr := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: blog2})

	// if updaterr != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("\nBlog has been updated!: %v", updateRes)

	// //deleting a new blog

	// DelRes, Delerr := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: res.GetBlog().GetId()})

	// if Delerr != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("\nBlog has been Deleted!: %v", DelRes)

	listStream, listErr := c.ListBlog(context.Background(), &blogpb.ListBlogRequest{})

	if listErr != nil {
		log.Fatal(err)
	}

	for {
		rcvRes, rcvErr := listStream.Recv()
		if rcvErr != nil {
			log.Fatalf("Error while receiving list of entries: %v", rcvErr)
		}

		fmt.Println(rcvRes.GetBlog())
	}

}
