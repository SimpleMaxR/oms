package types

import (
	"context"

	"github.com/simplemaxr/kitchen/services/common/genproto/orders"
)

// OrderService 接口，定义了订单服务需要实现的方法
type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
	GetOrder(context.Context) []*orders.Order
}
