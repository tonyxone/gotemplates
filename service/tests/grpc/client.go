package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	proto "gotemplates/api/grpc"
	"log"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	client := proto.NewUserClient(conn)
	resp, err := client.GetUserById(context.Background(), &proto.GetUserRequest{Id: uuid.New().String()})
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%v", resp)
}
