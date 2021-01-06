package main

import (
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	pb "rpcdemo/client_stream/proto"
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
	log.Println(Address + " net.Listing...")

	grpcServer := grpc.NewServer()
	pb.RegisterStreamClientServer(grpcServer, &SimpleService{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}

func (s *SimpleService) RouteList(srv pb.StreamClient_RouteListServer) error {
	for {
		res, err := srv.Recv()

		if err == io.EOF {
			return srv.SendAndClose(&pb.SimpleResponse{Value: "OK"})
		}

		if err != nil {
			return err
		}

		log.Println(res.StreamData)
	}
}
