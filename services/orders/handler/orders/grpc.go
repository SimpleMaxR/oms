package handler

import (
	"context"

	"github.com/simplemaxr/kitchen/services/common/genproto/orders"
	"github.com/simplemaxr/kitchen/services/orders/types"
	"google.golang.org/grpc"
)

type OrdersGrpcHandler struct {
	// service 注入
	orderService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrdersService(grpc *grpc.Server, orderService types.OrderService) {
	gRPCHandler := &OrdersGrpcHandler{
		orderService: orderService,
	}

	// 注册 OrderServiceServer
	orders.RegisterOrderServiceServer(grpc, gRPCHandler)
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	// handle 接口层逻辑

	order := &orders.Order{
		OrderId:    "001",
		CustomerId: req.CustomerId,
		ProductId:  req.ProductId,
		Quantity:   req.Quantity,
	}

	// 调用 service
	err := h.orderService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return &orders.CreateOrderResponse{
		Status: "ok",
	}, nil
}

func (h *OrdersGrpcHandler) GetOrder(ctx context.Context, req *orders.GetOrderRequest) (*orders.GetOrderResponse, error) {
	// handle 接口层逻辑

	o := h.orderService.GetOrder(ctx)

	res := &orders.GetOrderResponse{
		Orders: o,
	}

	return res, nil
}
