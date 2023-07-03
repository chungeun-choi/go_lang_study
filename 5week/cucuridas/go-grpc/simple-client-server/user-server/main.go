package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/cucuridas/go-grpc/data"
	userpb "github.com/cucuridas/go-grpc/protos/v1/user"
)


const portNumber = "9000"



type userServer struct {
	userpb.UserServer
}



func (s *userServer) GetUser(ctx context.Context,req *userpb.GetUserRequest)(*userpb.GetUserResponse,error){
	userID := req.UserId

	var userMessage *userpb.UserMessage

	for _,u := range(data.UserData) {
		if u.UserId != userID {
			continue
		}

		userMessage = u

		break
	}

	return &userpb.GetUserResponse{
		UserMessage: userMessage,
	},nil
}


func (s *userServer) ListUsers(ctx context.Context,req *userpb.ListUsersRequest)(*userpb.ListUsersResponse,error){
	userMessages := make([]*userpb.UserMessage, len(data.UserData))

	for i,u := range(data.UserData){
		userMessages[i] = u
	}

	return &userpb.ListUsersResponse{
		UserMessages: userMessages,
	},nil
}


func main() {
	lis,err := net.Listen("tcp", ":"+portNumber)

	if err != nil {
		log.Fatalf("Failed to liten %v", err)
		
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServer(grpcServer,&userServer{})

	log.Printf("Start gRPc server on %s port",portNumber)

	if err:= grpcServer.Serve(lis); err!= nil {
		log.Fatalf("failed to serve: %s",err)
	}

}
