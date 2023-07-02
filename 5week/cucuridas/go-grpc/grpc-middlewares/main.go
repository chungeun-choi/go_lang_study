package main

import (
	"context"
	"log"
	"net"

	"github.com/cucuridas/go-grpc/data"
	. "github.com/cucuridas/go-grpc/grpc-middlewares/meddle/auth"
	userpb "github.com/cucuridas/go-grpc/protos/v1/user"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


const portNumber = "9000"

type userServer struct {
	userpb.UserServer
}

func (s * userServer) GetUser(ctx context.Context,req *userpb.GetUserRequest)(*userpb.GetUserResponse,error){
	userID := req.UserId

	var userMessage *userpb.UserMessage

	for _, u := range data.UserData{
		if u.UserId != userID{
			continue
		}
		userMessage = u
		break
	}

	if userMessage == nil {
		return nil, status.Error(codes.NotFound, "user is no found")
	}

	return &userpb.GetUserResponse{
		UserMessage: userMessage,
	}, nil
}


func (s *userServer) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	userMessages := make([]*userpb.UserMessage, len(data.UserData))
	for i, u := range data.UserData {
		userMessages[i] = u
	}

	return &userpb.ListUsersResponse{
		UserMessages: userMessages,
	}, nil
}


func main() {
	lis, err := net.Listen("tcp",":"+portNumber)
	if err != nil {
		log.Fatalf("failed to listen: %v",err)
	}

	logrus.ErrorKey = "grpc.error"
	logrusEntry := logrus.NewEntry(logrus.StandardLogger())

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_auth.UnaryServerInterceptor(CustomAuthFunc),
			grpc_logrus.UnaryServerInterceptor(logrusEntry),
		)),
	)

	userpb.RegisterUserServer(grpcServer, &userServer{})

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}