package service

import (
	"context"

	"github.com/simplemaxr/kitchen/services/common/genproto/orders"
)

// ordersDb 是一个内存中的订单存储，用于模拟数据库操作
var ordersDb = make([]*orders.Order, 0)

// OrderService 结构体定义了订单服务
type OrderService struct {
	// 持久层接口放在这里
}

// NewOrderService 创建并返回一个新的 OrderService 实例
func NewOrderService() *OrderService {
	return &OrderService{}
}

// CreateOrder 创建一个新的订单并将其添加到 ordersDb 中
func (s *OrderService) CreateOrder(ctx context.Context, req *orders.Order) error {
	ordersDb = append(ordersDb, req)
	return nil
}

// GetOrder 返回所有存储在 ordersDb 中的订单
func (s *OrderService) GetOrder(ctx context.Context) []*orders.Order {
	return ordersDb
}
