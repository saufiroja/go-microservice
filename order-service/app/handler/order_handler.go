package handler

import (
	"context"
	"net/http"

	"github.com/saufiroja/order-service/app/client"
	"github.com/saufiroja/order-service/entity"
	"github.com/saufiroja/order-service/interfaces"
	"github.com/saufiroja/order-service/rpc/pb/order"
	pb "github.com/saufiroja/order-service/rpc/pb/order"
)

type OrderHandler struct {
	order.UnimplementedOrderServiceServer
	OrderService   interfaces.OrderService
	ProductService client.ProductServerClient
}

func NewOrderHandler(orderService interfaces.OrderService, client client.ProductServerClient) *OrderHandler {
	return &OrderHandler{
		OrderService:   orderService,
		ProductService: client,
	}
}

func (h *OrderHandler) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	product, err := h.ProductService.FindOne(req.ProductId)

	if err != nil {
		return &pb.CreateOrderResponse{Status: http.StatusBadRequest, Error: err.Error()}, nil
	} else if product.Status >= http.StatusNotFound {
		return &pb.CreateOrderResponse{Status: product.Status, Error: product.Error}, nil
	} else if int64(product.Data.Stock) < req.Quantity {
		return &pb.CreateOrderResponse{Status: http.StatusConflict, Error: "Stock too less"}, nil
	}

	order := entity.Order{
		Price:     int64(product.Data.Price),
		ProductID: product.Data.Id,
		UserID:    req.UserId,
	}

	_ = h.OrderService.CreateOrder(&order)

	res, err := h.ProductService.DecreaseStock(req.ProductId, order.ID)

	if err != nil {
		return &pb.CreateOrderResponse{Status: http.StatusBadRequest, Error: err.Error()}, nil
	} else if res.Status == http.StatusConflict {
		h.OrderService.DeleteOrder(order.ID)
		return &pb.CreateOrderResponse{Status: http.StatusConflict, Error: res.Error}, nil
	}

	return &pb.CreateOrderResponse{
		Status: http.StatusCreated,
		Id:     order.ID,
	}, nil
}
