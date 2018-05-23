package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/Graphmasters/presentations/grpc/03-client-streaming/proto"
	"io"
)

type server struct{}

func (server) Sum(stream pb.SumBuilder_SumServer) error {
	sum := 0.0
	for {
		x, err := stream.Recv()
		if err == io.EOF {
			// client finished sending, so we can return the response sentence
			return stream.SendAndClose(&pb.Response{
				Sum: sum,
			})
		}
		if err != nil {
			return err
		}

		// add number to sum
		log.Printf("adding number to sum: %f", x.X)
		sum += x.X
	}
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
	pb.RegisterSumBuilderServer(grpcServer, &server{})
	log.Printf("starting grpc server on %s", *addr)
	grpcServer.Serve(lis)
}
