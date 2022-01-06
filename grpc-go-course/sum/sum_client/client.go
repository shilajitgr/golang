package main

import (
	"context"
	"fmt"
	"log"

	"github.com/grpc-go-course/sum/sumpb"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}
	defer conn.Close()
	c := sumpb.NewSumServiceClient(conn)
	get_sum(c)
}

func get_sum(c sumpb.SumServiceClient) {
	fmt.Println("Starting Unary RPC...")
	req := &sumpb.SumRequest{
		Sum: &sumpb.Sum{
			Val1: 23,
			Val2: 54,
		},
	}
	res, err := c.Sum(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not get response %v", err)
	}

	log.Printf("Response recieved: %d ", res.GetResult())
}
