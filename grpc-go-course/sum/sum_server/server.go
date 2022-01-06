package main

import (
	"context"
	"log"
	"net"

	"github.com/grpc-go-course/sum/sumpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	a := req.GetSum().GetVal1()
	b := req.GetSum().GetVal2()
	result := a + b
	res := sumpb.SumResponse{
		Result: result,
	}

	return &res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Server could not listen %v", err)
	}

	s := grpc.NewServer()
	sumpb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Server could not start %v", err)
	}
}
