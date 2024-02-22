package main

import (
	"context"
	proto "github.com/tonggongzhiqiu/GoLearning/grpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type HelloService struct {
}

func (s *HelloService) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResp, error) {
	log.Println(req.Name)
	return &proto.HelloResp{Message: "hello, i am wufu"}, nil
}
func (s *HelloService) mustEmbedUnimplementedHelloServer() {

}

const (
	Address string = ":8000"
	Network string = "tcp"
)

func main() {
	listener, err := net.Listen(Network, Address)

	if err != nil {
		log.Panic("net.Listen err! err=%v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterHelloServer(grpcServer, &HelloService{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Panic("grpcServer.Serve err! err=%v", err)
	}
}
