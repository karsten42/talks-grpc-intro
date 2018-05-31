package main

import (
	"flag"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"

	"crypto/tls"
	"crypto/x509"
	pb "github.com/Graphmasters/presentations/grpc/07-mTLS/proto"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
)

var (
	server_crt = "./out/localhost.crt"
	server_key = "./out/localhost.key"
	server_ca  = "./out/CertAuth.crt"
)

type server struct {
}

func (server) Answer(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{
		Answer: "totally secure",
	}, nil
}

func newServer() *server {
	s := &server{}
	return s
}

func main() {
	addr := flag.String("addr", ":8000", "The address to bind the server to")

	// Load the certificates from disk
	certificate, err := tls.LoadX509KeyPair(server_crt, server_key)
	if err != nil {
		log.Fatalf("could not load server key pair: %s", err)
	}

	// Create a certificate pool from the certificate authority
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(server_ca)
	if err != nil {
		log.Fatalf("could not read ca certificate: %s", err)
	}

	// Append the client certificates from the CA
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("failed to append client certs")
	}

	// Create the TLS credentials
	creds := credentials.NewTLS(&tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
	})

	flag.Parse()
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterAnswerServiceServer(grpcServer, newServer())
	log.Printf("starting grpc server on %s", *addr)
	grpcServer.Serve(lis)
}
