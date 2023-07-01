package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)



const portNumber = "9000"



func main() {
	lis,err := net.Listen("tcp",":"+portNumber)

	if err !=nil {
		log.Fatalf("Failed to listen %v",err)
	}

	grpcServer := grpc.NewServer()

	log.Printf("Start gRPC server on %s port",portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Filed to serve: %s",err)
	}
}