package main

import (
	"context"
	"flag"
	pb "github.com/Graphmasters/presentations/grpc/01-single-call/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	serverAddr := flag.String("server_addr", "localhost:8000", "The server address in the format of host:port")

	flag.Parse()
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewAnswerServiceClient(conn)

	// ask for an answer
	answer, err := client.Answer(context.Background(), &pb.Request{
		Question: "what is this?",
	})
	if err != nil {
		log.Fatalf("failed to get answer: %v", err)
	}

	log.Printf("answer: %s", answer.Answer)
}
