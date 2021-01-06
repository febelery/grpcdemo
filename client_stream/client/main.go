package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "rpcdemo/client_stream/proto"
	"strconv"
)

const Address string = "localhost:8000"

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc dail err: %v", conn)
	}
	defer conn.Close()

	streamClient := pb.NewStreamClientClient(conn)
	stream, err := streamClient.RouteList(context.Background())
	if err != nil {
		log.Fatalf("client stream err: %v", err)
	}

	for n := 0; n < 10; n++ {
		err := stream.Send(&pb.StreamRequest{StreamData: "stream client rpc " + strconv.Itoa(n)})
		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("RouteList get response err: %v", err)
	}

	log.Println(res)
}
