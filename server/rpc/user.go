package rpc

import (
	"context"
	"github.com/hhandhuan/grpc-demo/proto"
)

type UserService struct{}

func (u *UserService) GetUser(ctx context.Context, request *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	return &proto.GetUserResponse{Id: request.Id}, nil
}