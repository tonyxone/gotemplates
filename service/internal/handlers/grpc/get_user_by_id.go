package grpc

import (
	"context"
	"fmt"
	"gotemplates/api/grpc"
)

func (h *UserServer) GetUserById(ctx context.Context, req *grpc.GetUserRequest) (*grpc.GetUserResponse, error) {
	fmt.Println("get user by id")
	return &grpc.GetUserResponse{
		Id:     req.GetId(),
		Name:   "Tao",
		Gender: "male",
	}, nil
}
