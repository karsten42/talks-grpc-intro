package main

import (
	"context"
	"flag"
	pb "github.com/Graphmasters/presentations/grpc/02-server-streaming/proto"
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

	// get N from cli
	nArg := flag.Arg(0)
	n, err := strconv.ParseInt(nArg, 10, 32)
	if err != nil {
		log.Fatalf("failed to parse %s as integer", nArg)
	}

	client := pb.NewFibonacciClient(conn)

	stream, err := client.Fib(context.Background(), &pb.Request{
		N: int32(n),
	})
	if err != nil {
		log.Fatalf("failed to get fib sequence: %v", err)
	}

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("Failed to receive a number : %v", err)
		}
		log.Printf("%d: %v", in.N, in.Res)
	}
}
