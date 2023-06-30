package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	gw "go-grpc/protos/v1/user" // Update
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9000", "gRPC server endpoint")
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterUserHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		glog.Fatal(err)
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	if err := http.ListenAndServe(":8081", mux); err != nil {
		glog.Fatal(err)
	}
}
