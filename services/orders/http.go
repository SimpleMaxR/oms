package main

import (
	"log"
	"net/http"

	handler "github.com/simplemaxr/kitchen/services/orders/handler/orders"
	"github.com/simplemaxr/kitchen/services/orders/service"
)

// httpServer 结构体，存有 addr
type httpServer struct {
	addr string
}

// 创建 httpServer，返回一个 httpServer 结构体的指针
func NewHttpServer(addr string) *httpServer {
	return &httpServer{
		addr: addr,
	}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	orderService := service.NewOrderService()
	orderHandler := handler.NewHttpOrderHandler(orderService)
	orderHandler.RegisterRouter(router)

	log.Println("Starting http server on ", s.addr)
	return http.ListenAndServe(s.addr, router)
}
