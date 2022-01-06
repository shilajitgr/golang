package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	greetManypb "github.com/grpc-go-course/greetMany/greetManypb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) GreetEveryone(stream greetManypb.GreetService_GreetEveryoneServer) error {
	log.Printf("GreetEveryone function was invoked...")
	var result string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Printf("Error while reading client stream %v", err)
			return err
		}

		firstName := req.GetGreetEveryone().GetFirstName()
		fmt.Printf("Received request from %v\n", firstName)
		result = "Hello " + firstName
		sendErr := stream.Send(&greetManypb.GreetEveryoneResponse{
			Result: result,
		})

		if sendErr != nil {
			log.Printf("Error while sending response to client %v", err)
			return err
		}

	}

	return nil
}

func (*server) LongGreet(stream greetManypb.GreetService_LongGreetServer) error {

	result := "Hello "
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetManypb.LongGreetResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Printf("Error while reading client stream %v", err)
			return err
		}

		firstName := req.GetLongGreet().GetFirstName()
		result += firstName + " "
	}
}

func (*server) GreetManyTimes(req *greetManypb.GreetManyTimesRequest, stream greetManypb.GreetService_GreetManyTimesServer) error {
	fmt.Println("Greet function was invoked with %v", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " Number " + strconv.Itoa(i)
		res := &greetManypb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Couldn't listen %v", err)
	}
	fmt.Println("Listening...")
	s := grpc.NewServer()
	greetManypb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}

}
