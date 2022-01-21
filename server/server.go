package server

import (
	"github.com/hhandhuan/grpc-demo/proto"
	"github.com/hhandhuan/grpc-demo/server/rpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

const PORT = "9001"

func Run() {
	server := grpc.NewServer()
	proto.RegisterSearchUserServiceServer(server, &rpc.SearchService{})
	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	server.Serve(lis)
}
