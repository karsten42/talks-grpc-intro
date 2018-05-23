package main

import (
	"context"
	"flag"
	pb "github.com/Graphmasters/presentations/grpc/04-bidirectional-streaming/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"strconv"
)

func main() {
	serverAddr := flag.String("server_addr", "localhost:8000", "The server address in the format of host:port")

	flag.Parse()
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewRunningAvgClient(conn)

	stream, err := client.Avg(context.Background())
	if err != nil {
		log.Fatalf("failed to get fib sequence: %v", err)
	}

	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a avg update : %v", err)
			}
			log.Printf("new avg: %f", in.Avg)
		}
	}()
	// iterate args
	for _, s := range flag.Args() {
		// parse number from commandline
		x, err := strconv.ParseFloat(s, 64)
		if err != nil {
			log.Fatalf("failed to parse input number %s: %v", s, err)
		}
		stream.Send(&pb.Request{
			X: x,
		})
	}

	// finish client sending
	stream.CloseSend()
	<-waitc

	log.Print("done")
}
