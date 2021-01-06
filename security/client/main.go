package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"rpcdemo/security/client/auth"
	pb "rpcdemo/security/proto"
)

const Address = "127.0.0.1:8000"

func main() {
	creds, err := credentials.NewClientTLSFromFile("./tls/server.pem", "")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}

	token := auth.Token{
		AppID:     "grpc_token",
		AppSecret: "grpc_secret",
	}
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&token))
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	grpcClient := pb.NewSimpleClient(conn)
	req := pb.SimpleRequest{Data: "grpc"}

	res, err := grpcClient.Route(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}

	log.Println(res)
}
