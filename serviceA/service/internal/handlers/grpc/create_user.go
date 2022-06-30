package grpc

import (
	"context"
	"serviceA/api/grpc"
)

func (h *UserServer) CreateUser(context.Context, *grpc.CreateUserRequest) (*grpc.CreateUserResponse, error) {
	return nil, nil
}
