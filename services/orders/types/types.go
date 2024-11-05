package types

import (
	"context"

	"github.com/simplemaxr/kitchen/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
	GetOrder(context.Context) []*orders.Order
}
