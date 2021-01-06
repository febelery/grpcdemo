package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
	pb "rpcdemo/security/proto"
)

const (
	Address string = ":8000"
	Network string = "tcp"
)

type SimpleService struct{}

func main() {
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile("./tls/server.pem", "./tls/server.key")
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}

	// 拦截器
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if err = Check(ctx); err != nil {
			log.Printf("grpc interceptor err: %v", err)
			return
		}

		return handler(ctx, req)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(interceptor))
	pb.RegisterSimpleServer(grpcServer, &SimpleService{})
	log.Println(Address + " net.Listing whth TLS and token...")

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}

func (s *SimpleService) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	res := pb.SimpleResponse{
		Code:  200,
		Value: "hello security " + req.Data,
	}

	return &res, nil
}

func Check(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "get token failed")
	}

	// 一些校验
	log.Println(md)

	return nil
}
