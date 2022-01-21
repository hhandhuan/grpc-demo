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

	client := proto.NewSearchUserServiceClient(conn)
	resp, err := client.SearchUserDetail(context.Background(), &proto.SearchUserDetailRequest{
		Id: 22,
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(resp.GetUser().GetId())
}