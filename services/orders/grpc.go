package main

import (
	"log"
	"net"

	handler "github.com/simplemaxr/kitchen/services/orders/handler/orders"
	"github.com/simplemaxr/kitchen/services/orders/service"
	"google.golang.org/grpc"
)

// gRPCServer 结构体用于封装 gRPC 服务器的配置
type gRPCServer struct {
	addr string // 服务器监听地址
}

// NewGRPCServer 创建一个新的 gRPCServer 实例
func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{
		addr: addr,
	}
}

// Run 启动 gRPC 服务器
func (s *gRPCServer) Run() error {
	// 创建 TCP 监听器
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	// 创建 gRPC 服务器实例
	gRPCServer := grpc.NewServer()

	// 创建并注册订单服务
	orderService := service.NewOrderService()
	handler.NewGrpcOrdersService(gRPCServer, orderService)

	// 输出服务器启动信息
	log.Println("gRPC server started on", s.addr)

	// 启动 gRPC 服务器
	return gRPCServer.Serve(lis)
}
