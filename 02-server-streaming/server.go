package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/Graphmasters/presentations/grpc/02-server-streaming/proto"
	"time"
)

type server struct{}

func (server) Fib(req *pb.Request, stream pb.Fibonacci_FibServer) error {
	a := uint64(0)
	b := uint64(1)
	stream.Send(&pb.Response{
		N:   int32(1),
		Res: b,
	})

	for i := 2; i <= int(req.N); i++ {
		bb := a + b
		stream.Send(&pb.Response{
			N:   int32(i),
			Res: bb,
		})
		a = b
		b = bb
		time.Sleep(10 * time.Millisecond)
	}
	return nil
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
	pb.RegisterFibonacciServer(grpcServer, &server{})
	log.Printf("starting grpc server on %s", *addr)
	grpcServer.Serve(lis)
}
