package server

import (
	"context"
	"github.com/hhandhuan/grpc-demo/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type HelloService struct{}

func (i *HelloService) GetUser(ctx context.Context, request *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	log.Println(request.Id)
	return &proto.GetUserResponse{Id: 100}, nil
}

const PORT = "9001"

func Run() {
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &HelloService{})
	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	server.Serve(lis)
}
