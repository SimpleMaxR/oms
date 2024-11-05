package main

import (
	"log"
	"net"

	handler "github.com/simplemaxr/kitchen/services/orders/handler/orders"
	"github.com/simplemaxr/kitchen/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{
		addr: addr,
	}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	gRPCServer := grpc.NewServer()

	// 注册服务
	orderService := service.NewOrderService()
	handler.NewGrpcOrdersService(gRPCServer, orderService)

	log.Println("gRPC server started on ", s.addr)
	return gRPCServer.Serve(lis)
}
