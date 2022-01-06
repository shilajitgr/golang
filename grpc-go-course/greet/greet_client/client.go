package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/grpc-go-course/greet/greetpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

func main() {
	creds, sslErr := credentials.NewClientTLSFromFile("ssl/ca.crt", "") // passing path to cert authority trust certificate

	if sslErr != nil {
		log.Fatalf("Failed while loading SSL certs %v", sslErr)
		return
	}

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatalf("Couldn't connect to server %v", err)
	}
	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)
	doUnary(c)
	// doUnaryWithDeadline(c, 5)
	// doUnaryWithDeadline(c, 1)
}

func doUnaryWithDeadline(c greetpb.GreetServiceClient, seconds time.Duration) {
	fmt.Printf("Starting Unary RPC with %v timeout...", seconds)
	req := &greetpb.DeadlineRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Shilajit",
			LastName:  "Roy",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), seconds*1000*time.Millisecond)
	// context.withDeadline(x,y), where y is the time object which specifies the time by which the operation should return
	// either successfully or with deadline exceeded error
	defer cancel()

	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("\nTimeout was exceeded!")
			} else {
				fmt.Printf("Unexpected Error: %v", statusErr.Details())
			}
		}
		log.Fatalf("Error while calling greetWithDeadline %v", err)
		return
	}

	log.Printf("\nResponse from greetWithDeadline %v", res.GetResult())

}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Shilajit",
			LastName:  "Roy",
		},
	}

	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling greet %v", err)
	}

	log.Printf("Response from Greet %v", res)

}
