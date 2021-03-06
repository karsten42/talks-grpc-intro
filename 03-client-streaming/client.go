package main

import (
	"context"
	"flag"
	"log"
	"strconv"
	"time"

	pb "github.com/Graphmasters/presentations/grpc/03-client-streaming/proto"
	"google.golang.org/grpc"
)

func main() {
	serverAddr := flag.String("server_addr", "localhost:8000", "The server address in the format of host:port")

	flag.Parse()
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewSumBuilderClient(conn)

	stream, err := client.Sum(context.Background())
	if err != nil {
		log.Fatalf("failed to get fib sequence: %v", err)
	}

	// iterate args
	for _, s := range flag.Args() {
		// parse number from commandline
		x, err := strconv.ParseFloat(s, 64)
		if err != nil {
			log.Fatalf("failed to parse input number %s: %v", s, err)
		}
		log.Printf("sending number: %f", x)
		if err := stream.Send(&pb.Request{
			X: x,
		}); err != nil {
			log.Fatalf("boom: %v", err)
		}
		time.Sleep(time.Second)
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to build sentence: %v", err)
	}

	log.Printf("sum: %f", resp.Sum)
}
