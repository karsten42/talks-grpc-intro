package main

import (
	"flag"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/Graphmasters/presentations/grpc/01-single-call/proto"
)

type server struct {
}

func (server) Answer(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{
		Answer: "awesome",
	}, nil
}

func newServer() *server {
	s := &server{}
	return s
}

func main() {
	addr := flag.String("addr", ":8000", "The address to bind the server to")

	flag.Parse()

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterAnswerServiceServer(grpcServer, newServer())
	log.Printf("starting grpc server on %s", *addr)
	grpcServer.Serve(lis)
}
