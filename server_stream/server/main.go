package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	pb "rpcdemo/server_stream/proto"
	"strconv"
)

const (
	Address string = ":8000"
	Network string = "tcp"
)

type StreamService struct{}

func main() {
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println(Address + " net.Listing...")

	grpcServer := grpc.NewServer()
	pb.RegisterStreamServerServer(grpcServer, &StreamService{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}

func (s *StreamService) ListValue(req *pb.SimpleRequest, srv pb.StreamServer_ListValueServer) error {
	for n := 0; n < 10; n++ {
		err := srv.Send(&pb.StreamResponse{
			StreamValue: req.Data + strconv.Itoa(n),
		})

		if err != nil {
			return err
		}
	}

	return nil
}
