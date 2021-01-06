package cred

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func TLSInterceptor() grpc.ServerOption {
	creds, err := credentials.NewServerTLSFromFile("./tls/server.pem", "./tls/server.key")
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}

	return grpc.Creds(creds)
}
