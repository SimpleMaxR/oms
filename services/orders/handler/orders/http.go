package handler

import (
	"net/http"

	"github.com/simplemaxr/kitchen/services/common/genproto/orders"
	"github.com/simplemaxr/kitchen/services/common/util"
	"github.com/simplemaxr/kitchen/services/orders/types"
)

type OrderHttpHandler struct {
	ordersService types.OrderService
}

func NewHttpOrderHandler(service types.OrderService) *OrderHttpHandler {
	handler := &OrderHttpHandler{
		ordersService: service,
	}
	return handler
}

func (h *OrderHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrderHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	err := util.ParseJSON(r, req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	order := &orders.Order{
		OrderId:    "001",
		CustomerId: req.CustomerId,
		ProductId:  req.ProductId,
		Quantity:   req.Quantity,
	}

	err = h.ordersService.CreateOrder(r.Context(), order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := &orders.CreateOrderResponse{
		Status: "ok",
	}
	util.WriteJSON(w, http.StatusOK, res)
}
