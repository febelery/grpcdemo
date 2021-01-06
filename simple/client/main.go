package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "rpcdemo/simple/proto"
	"time"
)

const Address string = "localhost:8000"

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	grpcClient := pb.NewSimpleClient(conn)
	req := pb.SimpleRequest{Data: "grpc"}

	ctx := context.Background()
	clientDeadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(ctx, clientDeadline)
	defer cancel()

	res, err := grpcClient.Route(ctx, &req)
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}

	log.Println(res)
}
