package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	pb "github.com/Graphmasters/presentations/grpc/07-mTLS/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

var (
	client_crt = "./out/localhost.crt"
	client_key = "./out/localhost.key"
	client_ca  = "./out/CertAuth.crt"
)

func main() {
	serverAddr := flag.String("server_addr", "localhost:8000", "The server address in the format of host:port")
	flag.Parse()

	// Load the client certificates from disk
	certificate, err := tls.LoadX509KeyPair(client_crt, client_key)
	if err != nil {
		log.Fatalf("could not load client key pair: %s", err)
	}

	// Create a certificate pool from the certificate authority
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(client_ca)
	if err != nil {
		log.Fatalf("could not read ca certificate: %s", err)
	}

	// Append the certificates from the CA
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append ca certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		ServerName:   "localhost", // NOTE: this is required!
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
	})
	println(creds)

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
