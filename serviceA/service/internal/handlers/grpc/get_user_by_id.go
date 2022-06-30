package grpc

import (
	"context"
	"serviceA/api/grpc"
)

func (h *UserServer) GetUserById(ctx context.Context, req *grpc.GetUserRequest) (*grpc.GetUserResponse, error) {
	return nil, nil
}
