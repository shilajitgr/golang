package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/grpc-go-course/greetMany/greetManypb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Couldn't connect to server %v", err)
	}
	defer conn.Close()
	c := greetManypb.NewGreetServiceClient(conn)
	// doStream(c)
	// doClientStreaming(c)
	doBidiStreaming(c)
}

func doBidiStreaming(c greetManypb.GreetServiceClient) {
	fmt.Println("Starting bi-directional streaming...")
	//create stream by invoking client
	resStream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while calling greetManyTimes %v", err)
		return
	}

	requests := []*greetManypb.GreetEveryoneRequest{
		&greetManypb.GreetEveryoneRequest{
			GreetEveryone: &greetManypb.Greeting{
				FirstName: "Shilajit",
				LastName:  "Roy",
			},
		},
		&greetManypb.GreetEveryoneRequest{
			GreetEveryone: &greetManypb.Greeting{
				FirstName: "Pakshal",
				LastName:  "B",
			},
		},
		&greetManypb.GreetEveryoneRequest{
			GreetEveryone: &greetManypb.Greeting{
				FirstName: "Hitesh",
				LastName:  "K B",
			},
		},
	}

	waitc := make(chan string)
	//send messages to the client(using go routine)
	go func() {
		for _, req := range requests {
			fmt.Printf("Sending msg:%v\n", req)
			resStream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		resStream.CloseSend()
	}()

	//recv messages from the client(using go routine)
	go func() {

		for {
			resp, err := resStream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving greetEveryone response %v", err)
				break
			}
			log.Printf("Response from GreetEveryone %v\n", resp.GetResult())
		}
		close(waitc) // this will close the channel and thereby ending program execution

	}()

	//block until everything is done
	<-waitc // this statement is trying to extract something from a channel, but no msg has been placed into the
	// the channel, hence this statement acts as a blocking statement
}

func doStream(c greetManypb.GreetServiceClient) {
	fmt.Println("Starting server streaming RPC...")
	req := &greetManypb.GreetManyTimesRequest{
		Greeting: &greetManypb.Greeting{
			FirstName: "Shilajit",
			LastName:  "Roy",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling greetManyTimes %v", err)
	}

	for {
		msg, err := resStream.Recv() // when stream response ends we will recv an io endoffile error
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while receiving greetManyTimes response %v", err)
		}
		log.Printf("Response from GreetManyTimes %v", msg.GetResult())
	}

}

func doClientStreaming(c greetManypb.GreetServiceClient) {
	reqStream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet %v", err)
	}

	requests := []*greetManypb.LongGreetRequest{
		&greetManypb.LongGreetRequest{
			LongGreet: &greetManypb.Greeting{
				FirstName: "Shilajit",
				LastName:  "Roy",
			},
		},
		&greetManypb.LongGreetRequest{
			LongGreet: &greetManypb.Greeting{
				FirstName: "Pakshal",
				LastName:  "B",
			},
		},
		&greetManypb.LongGreetRequest{
			LongGreet: &greetManypb.Greeting{
				FirstName: "Hitesh",
				LastName:  "K B",
			},
		},
	}
	for req := range requests {
		reqStream.Send(requests[req])
		time.Sleep(time.Millisecond * 1000)
		fmt.Printf("Sent request : %v", requests[req])
	}

	res, err := reqStream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving LongGreet response %v", err)
	}

	fmt.Printf("Response rcvd: %v", res)

}
