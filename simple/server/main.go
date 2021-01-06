package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	pb "rpcdemo/simple/proto"
	"runtime"
	"time"
)

type SimpleService struct{}

const (
	Address string = ":8000"
	Network string = "tcp"
)

func main() {
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println(Address + " net.Listening ...")

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleServer(grpcServer, &SimpleService{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}

func (s *SimpleService) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	data := make(chan *pb.SimpleResponse, 1)
	go func(ctx context.Context, req *pb.SimpleRequest, data chan<- *pb.SimpleResponse) {
		select {
		case <-ctx.Done():
			log.Println(ctx.Err())
			runtime.Goexit() //超时后退出该Go协程
		case <-time.After(2 * time.Second): // 模拟耗时操作
			res := pb.SimpleResponse{
				Code:  200,
				Value: "hello " + req.Data,
			}

			data <- &res
		}
	}(ctx, req, data)

	select {
	case res := <-data:
		return res, nil
	case <-ctx.Done():
		return nil, status.Errorf(codes.Canceled, "Client cancelled, abandoning.")
	}
}
