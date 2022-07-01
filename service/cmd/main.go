package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	handler "serviceA/service/internal/handlers/grpc"

	proto "serviceA/api/grpc"
)

func main() {
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})

	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
