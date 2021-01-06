package gateway

import (
	"context"
	"crypto/tls"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"io/ioutil"
	"log"
	"net/http"
	pb "rpcdemo/demo/proto"
	"strings"
)

func ProvideHTTP(endpoint string, grpcServer *grpc.Server) *http.Server {
	ctx := context.Background()

	creds, err := credentials.NewClientTLSFromFile("./tls/server.pem", "")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}

	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	// grpc-gateway的请求复用器。它将http请求与模式匹配，并调用相应的处理程序。
	gwmux := runtime.NewServeMux()
	err = pb.RegisterSimpleHandlerFromEndpoint(ctx, gwmux, endpoint, opts)
	if err != nil {
		log.Fatalf("Register Endpoint err: %v", err)
	}

	// http的请求复用器
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	log.Println(endpoint + " http.Listing whth TLS and token...")

	return &http.Server{
		Addr:      endpoint,
		Handler:   grpcHandlerFunc(grpcServer, mux),
		TLSConfig: getTlsConfig(),
	}
}

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

func getTlsConfig() *tls.Config {
	cert, _ := ioutil.ReadFile("./tls/server.pem")
	key, _ := ioutil.ReadFile("./tls/server.key")

	pair, err := tls.X509KeyPair(cert, key)
	if err != nil {
		grpclog.Fatalf("TLS KeyPair err: %v\n", err)
	}

	return &tls.Config{
		Certificates: []tls.Certificate{pair},
		NextProtos:   []string{http2.NextProtoTLS}, // HTTP2 TLS支持
	}
}
