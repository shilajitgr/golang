package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	primeDecomposepb "github.com/grpc-go-course/PrimeDecomposeStream/primepb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}
	defer conn.Close()
	c := primeDecomposepb.NewPrimeServiceClient(conn)
	// get_factor(c)
	// get_avg(c)
	// get_max(c)
	get_sqrt(c)
}

func get_sqrt(c primeDecomposepb.PrimeServiceClient) {
	fmt.Println("Starting a unary RPC with error handling")
	// error request
	res, err := c.SquareRoot(context.Background(), &primeDecomposepb.SqrtRequest{
		Val: -4,
	})
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("A negative number was sent as argument!")
			}
		} else {
			log.Fatalf("Big error calling SquareRoot: %v", err)
		}
		return
	}
	log.Printf("Response received from SquareRoot: %v", res.GetSqrt())
}

func get_max(c primeDecomposepb.PrimeServiceClient) {

	fmt.Println("Starting client streaming RPC...")

	reqStream, err := c.FindMax(context.Background())

	if err != nil {
		log.Fatalf("Error while connecting to server, %v", err)
		return
	}

	vals := []int32{2, 5, 3, 4, 6, 1}
	waitc := make(chan string)
	go func() {
		for _, val := range vals {
			err := reqStream.Send(&primeDecomposepb.FindMaxRequest{
				Val: val,
			})
			log.Printf("Request sent: %d\n", val)

			if err != nil {
				log.Fatalf("Error while sending requests to server, %v", err)
				break
			}
			time.Sleep(1000 * time.Millisecond)
		}
		reqStream.CloseSend()
	}()

	go func() {
		for {
			resp, err := reqStream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while receiving response from server, %v", err)
				break
			}
			log.Printf("Response from server: %d\n\n", resp.GetMax())
		}
		close(waitc)
	}()

	<-waitc
}

func get_avg(c primeDecomposepb.PrimeServiceClient) {
	fmt.Println("Starting client streaming RPC...")

	inputStream, err := c.ComputeAvg(context.Background())

	if err != nil {
		log.Fatalf("Error while calling ComputeAvg %v", err)
	}

	requests := []*primeDecomposepb.AvgRequest{
		&primeDecomposepb.AvgRequest{
			Input: &primeDecomposepb.Number{
				Val: 3,
			},
		},
		&primeDecomposepb.AvgRequest{
			Input: &primeDecomposepb.Number{
				Val: 5,
			},
		},
		&primeDecomposepb.AvgRequest{
			Input: &primeDecomposepb.Number{
				Val: 4,
			},
		},
		&primeDecomposepb.AvgRequest{
			Input: &primeDecomposepb.Number{
				Val: 7,
			},
		},
	}

	for _, req := range requests {
		inputStream.Send(req)
		time.Sleep(time.Millisecond * 1000)
	}
	res, err := inputStream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while recieving ComputeAvg response %v", err)
	}

	log.Printf("Response received: %f", res.GetAvg())

}

func get_factor(c primeDecomposepb.PrimeServiceClient) {
	fmt.Println("Starting Unary RPC...")
	req := &primeDecomposepb.PrimeRequest{
		Input: &primeDecomposepb.Number{
			Val: 120,
		},
	}
	factorStream, err := c.Number(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not get response %v", err)
	}
	// var factor *primeDecomposepb.PrimeResponse
	for {
		factor, err := factorStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while receiving Prime factors %v", err)
		}
		log.Printf("Prime factor received %d", factor.GetPrimeFactor())
	}

}
