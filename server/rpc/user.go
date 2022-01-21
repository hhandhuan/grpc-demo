package rpc

import (
	"context"
	"errors"
	"github.com/hhandhuan/grpc-demo/proto"
)

type SearchService struct{}

func (u *SearchService) SearchUserDetail(ctx context.Context, request *proto.SearchUserDetailRequest) (*proto.SearchUserDetailResponse, error) {
	if request.Id == 22 {
		return nil, errors.New("ID error")
	}
	return &proto.SearchUserDetailResponse{
		User: &proto.User{
			Id:       11,
			Username: "zhangsan",
			Nickname: "resh",
			State:    1,
		},
	}, nil
}
