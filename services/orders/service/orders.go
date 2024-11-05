package service

import (
	"context"

	"github.com/simplemaxr/kitchen/services/common/genproto/orders"
)

var ordersDb = make([]*orders.Order, 0)

type OrderService struct {
	// 持久层接口放在这里
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *orders.Order) error {
	ordersDb = append(ordersDb, req)
	return nil
}

func (s *OrderService) GetOrder(ctx context.Context) []*orders.Order {
	return ordersDb
}
