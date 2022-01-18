package main

import (
	"context"
	"github.com/hhandhuan/grpc-demo/proto"
	"log"

	"google.golang.org/grpc"

)

const PORT = "9001"

func main() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := proto.NewUserClient(conn)
	resp, err := client.GetUser(context.Background(), &proto.GetUserRequest{
		Id: 22,
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	log.Printf("resp: %s", resp.GetId())
}