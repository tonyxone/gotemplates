package grpc

import "serviceA/api/grpc"

type UserServer struct {
	grpc.UnimplementedUserServer
}
