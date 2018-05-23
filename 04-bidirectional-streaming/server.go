package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/Graphmasters/presentations/grpc/04-bidirectional-streaming/proto"
	"io"
)

type server struct{}

func (server) Avg(stream pb.RunningAvg_AvgServer) error {
	total := 0.0
	cnt := 0
	for {
		x, err := stream.Recv()
		if err == io.EOF {
			// finished
			return nil
		}
		if err != nil {
			return err
		}

		// add number to sum
		log.Printf("adding number to avg: %f", x.X)
		total += x.X
		cnt++

		// send current avg back to client
		stream.Send(&pb.Response{
			Avg: total / float64(cnt),
		})
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
	pb.RegisterRunningAvgServer(grpcServer, &server{})
	log.Printf("starting grpc server on %s", *addr)
	grpcServer.Serve(lis)
}
