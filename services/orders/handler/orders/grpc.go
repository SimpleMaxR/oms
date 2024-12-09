package handler

import (
	"context"

	"github.com/simplemaxr/kitchen/services/common/genproto/orders"
	"github.com/simplemaxr/kitchen/services/orders/types"
	"google.golang.org/grpc"
)

// OrdersGrpcHandler 处理 gRPC 订单服务请求
type OrdersGrpcHandler struct {
	orderService types.OrderService // 注入订单服务的实现
	orders.UnimplementedOrderServiceServer
}

// NewGrpcOrdersService 创建并注册 gRPC 订单服务
func NewGrpcOrdersService(grpc *grpc.Server, orderService types.OrderService) {
	gRPCHandler := &OrdersGrpcHandler{
		orderService: orderService,
	}
	orders.RegisterOrderServiceServer(grpc, gRPCHandler) // 使用 grpc Server 实例和 OrdersGrpcHandler 实例注册服务
}

// CreateOrder 处理创建订单的 gRPC 请求
func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderId:    "001", // 注意：这里应该使用动态生成的订单ID
		CustomerId: req.CustomerId,
		ProductId:  req.ProductId,
		Quantity:   req.Quantity,
	}

	err := h.orderService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return &orders.CreateOrderResponse{
		Status: "ok",
	}, nil
}

// GetOrder 处理获取订单的 gRPC 请求
func (h *OrdersGrpcHandler) GetOrder(ctx context.Context, req *orders.GetOrderRequest) (*orders.GetOrderResponse, error) {
	// handle 接口层逻辑

	o := h.orderService.GetOrder(ctx)

	res := &orders.GetOrderResponse{
		Orders: o,
	}

	return res, nil
}
