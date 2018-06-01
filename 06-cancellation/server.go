package main

import (
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"

	"time"

	pb "github.com/Graphmasters/presentations/grpc/02-server-streaming/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

func (server) Fib(req *pb.Request, stream pb.Fibonacci_FibServer) error {
	ctx := stream.Context()
	a := uint64(0)
	b := uint64(1)
	stream.Send(&pb.Response{
		N:   int32(1),
		Res: b,
	})

	for i := 2; i <= int(req.N); i++ {
		if ctx.Err() != nil {
			log.Print("call canceled")
			return status.New(codes.Canceled, "Client cancelled, abandoning.").Err()
		}

		bb := a + b
		stream.Send(&pb.Response{
			N:   int32(i),
			Res: bb,
		})
		a = b
		b = bb
		time.Sleep(1 * time.Second)
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
