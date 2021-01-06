package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"rpcdemo/demo/client/auth"
	pb "rpcdemo/demo/proto"
)

const Address = "127.0.0.1:8000"

func main() {
	creds, err := credentials.NewClientTLSFromFile("./tls/server.pem", "")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}

	token := auth.Token{
		Value: "bearer grpc.auth.token",
	}

	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&token))
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	grpcClient := pb.NewSimpleClient(conn)
	req := pb.SimpleRequest{
		Data: "grpc",
		Do:   2,
	}

	res, err := grpcClient.Route(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}

	log.Println(res)
}
