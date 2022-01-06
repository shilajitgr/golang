package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"

	primeDecomposepb "github.com/grpc-go-course/PrimeDecomposeStream/primepb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func checkPrime(k int32) bool {
	var i int32
	for i = 2; i <= k/2; i++ {
		if k%i == 0 {
			return false
		}
	}
	return true
}

func (*server) Number(req *primeDecomposepb.PrimeRequest, stream primeDecomposepb.PrimeService_NumberServer) error {

	input := req.GetInput().GetVal()
	log.Printf("Prime factors of %d---\n", input)
	var k int32 = 2
	for {
		if input%k == 0 {
			log.Printf("%d\n", k)
			if checkPrime(k) {
				res := &primeDecomposepb.PrimeResponse{
					PrimeFactor: k,
				}
				stream.Send(res)
				input = input / k
			} else {
				k = k + 1
			}

		} else {
			k = k + 1
		}

		if input == 0 || k > input {
			break
		}
	}
	return nil
}

func (*server) ComputeAvg(stream primeDecomposepb.PrimeService_ComputeAvgServer) error {
	var count float32
	var result float32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&primeDecomposepb.AvgResponse{
				Avg: result / count,
			})
		}

		if err != nil {
			log.Printf("Error while reading client stream %v", err)
			return err
		}

		num := req.GetInput().GetVal()
		result += float32(num)
		count++
	}
}

func (*server) FindMax(stream primeDecomposepb.PrimeService_FindMaxServer) error {
	fmt.Println("FindMax Service invoked...")
	var i int16
	var max int32
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Printf("Error while reading client stream %v", err)
			return err
		}

		num := req.GetVal()
		if i == 0 {
			max = num
			i += 1
		} else {
			if num > max {
				max = num
			}
		}

		err = stream.Send(&primeDecomposepb.FindMaxResponse{
			Max: max,
		})

		if err != nil {
			log.Printf("Error while sending response %v", err)
			return err
		}
	}
}

func (*server) SquareRoot(ctx context.Context, in *primeDecomposepb.SqrtRequest) (*primeDecomposepb.SqrtResponse, error) {

	val := in.GetVal()

	if val < 0 {
		return nil, grpc.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received negative value : %v", val),
		)
	}

	return &primeDecomposepb.SqrtResponse{
		Sqrt: float32(math.Sqrt(float64(val))),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Server could not listen %v", err)
	}

	s := grpc.NewServer()
	primeDecomposepb.RegisterPrimeServiceServer(s, &server{})

	// register reflection services on the server
	reflection.Register(s) // now start evansc cli with the cmd(from cmd): evans -p [port num] -r , here r is for reflection

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Server could not start %v", err)
	}
}
