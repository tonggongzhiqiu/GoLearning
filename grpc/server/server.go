package main

import (
	proto "GoLearning/grpc/proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type HelloService struct {
	proto.UnimplementedHelloServer
}

func (s *HelloService) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResp, error) {
	log.Println(req.Name)
	return &proto.HelloResp{Message: "hello, i am " + req.GetName()}, nil
}

const (
	Address string = ":8000"
	Network string = "tcp"
)

func main() {
	listener, err := net.Listen(Network, Address)
	log.Println("start listen...")

	if err != nil {
		log.Panic("net.Listen err! err=%v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterHelloServer(grpcServer, &HelloService{})
	log.Println("register helloServer success")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Panic("grpcServer.Serve err! err=%v", err)
	}
}
