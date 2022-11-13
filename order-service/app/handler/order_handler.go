package handler

import (
	"context"
	"net/http"

	"github.com/saufiroja/order-service/entity"
	"github.com/saufiroja/order-service/interfaces"
	"github.com/saufiroja/order-service/rpc/pb/order"
)

type OrderHandler struct {
	order.UnimplementedOrderServiceServer
	OrderService interfaces.OrderService
}

func NewOrderHandler(orderService interfaces.OrderService) *OrderHandler {
	return &OrderHandler{
		OrderService: orderService,
	}
}

func (h *OrderHandler) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	data := &entity.Order{
		ProductID: req.ProductId,
		UserID:    req.UserId,
		Quantity:  req.Quantity,
	}

	err := h.OrderService.CreateOrder(data)
	if err != nil {
		return nil, err
	}

	return &order.CreateOrderResponse{
		Id:     data.ID,
		Status: http.StatusCreated,
	}, nil
}
