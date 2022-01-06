package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/grpc-go-course/greet/greetpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("Greet function was invoked with %v", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := greetpb.GreetResponse{
		Result: result,
	}
	return &res, nil
}

func (*server) GreetWithDeadline(ctx context.Context, req *greetpb.DeadlineRequest) (*greetpb.DeadlineResponse, error) {
	fmt.Println("GreetWithDeadline function was invoked with %v", req)
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			// client cancelled the request
			return nil, status.Error(codes.Canceled, "Client cancelled request!")
		}
		fmt.Println("sleeping...")
		time.Sleep(1 * time.Second)
	}
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := greetpb.DeadlineResponse{
		Result: result,
	}
	return &res, nil
}

func main() {

	creds, sslErr := credentials.NewServerTLSFromFile("ssl/server.crt", "ssl/server.pem")

	if sslErr != nil {
		log.Fatalf("Failed while loading SSL certs %v", sslErr)
		return
	}

	s := grpc.NewServer(grpc.Creds(creds))

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Couldn't listen %v", err)
	}
	fmt.Println("Listening...")
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}

}
