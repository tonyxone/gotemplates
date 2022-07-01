package grpc

import (
	"context"
	"gotemplates/api/grpc"
)

func (h *UserServer) CreateUser(ctx context.Context, req *grpc.CreateUserRequest) (*grpc.CreateUserResponse, error) {
	return &grpc.CreateUserResponse{
		Id: req.GetId(),
	}, nil
}
