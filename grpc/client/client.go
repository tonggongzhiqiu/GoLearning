package main

import (
	proto "GoLearning/grpc/proto"
	"context"
	"google.golang.org/grpc"
	"log"
)

const (
	ServerAddress string = ":8000"
)

func main() {
	conn, err := grpc.Dial(ServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("net.Connect err: %v", err)
	}
	defer conn.Close()

	grpcClient := proto.NewHelloClient(conn)
	req := proto.HelloRequest{
		Name: "wufu",
	}

	res, err := grpcClient.SayHello(context.Background(), &req)
	if err != nil {
		log.Fatal("call sayHello err: %v", err)
	}
	log.Println(res)
}
