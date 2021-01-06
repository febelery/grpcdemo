package main

import (
	"context"
	"crypto/tls"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcauth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcvalidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "rpcdemo/demo/proto"
	"rpcdemo/demo/server/gateway"
	"rpcdemo/demo/server/middleware/auth"
	"rpcdemo/demo/server/middleware/cred"
	"rpcdemo/demo/server/middleware/recovery"
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

	grpcServer := grpc.NewServer(cred.TLSInterceptor(),
		grpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer(
			grpcvalidator.UnaryServerInterceptor(),
			grpcauth.UnaryServerInterceptor(auth.Interceptor),
			grpcrecovery.UnaryServerInterceptor(recovery.Interceptor()),
		)),
		grpc.StreamInterceptor(grpcmiddleware.ChainStreamServer(
			grpcvalidator.StreamServerInterceptor(),
			grpcauth.StreamServerInterceptor(auth.Interceptor),
			grpcrecovery.StreamServerInterceptor(recovery.Interceptor()),
		)),
	)

	pb.RegisterSimpleServer(grpcServer, &SimpleService{})
	log.Println(Address + " net.Listing whth TLS and token...")

	httpServer := gateway.ProvideHTTP(Address, grpcServer)

	if err = httpServer.Serve(tls.NewListener(listener, httpServer.TLSConfig)); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func (s *SimpleService) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	res := pb.SimpleResponse{
		Code:  200,
		Value: "hello " + req.Data,
	}

	return &res, nil
}
