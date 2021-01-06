package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	pb "rpcdemo/server_stream/proto"
)

const Address string = "localhost:8000"

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	grpcClient := pb.NewStreamServerClient(conn)
	req := pb.SimpleRequest{Data: "stream server grpc"}

	stream, err := grpcClient.ListValue(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call ListStr err: %v", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("ListStr get stream err: %v", err)
		}

		log.Println(res.StreamValue)
	}

	// 可以使用CloseSend()关闭stream，这样服务端就不会继续产生流消息
	// 调用CloseSend()后，若继续调用Recv()，会重新激活stream，接着之前结果获取消息
	// stream.CloseSend()
}
