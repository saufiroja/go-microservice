package service

import (
	"github.com/saufiroja/order-service/app/client"
	"github.com/saufiroja/order-service/entity"
	"github.com/saufiroja/order-service/interfaces"
)

type OrderService struct {
	OrderRepository interfaces.OrderRepository
	ProductService  client.ProductServerClient
}

func NewOrderService(orderRepository interfaces.OrderRepository, productService client.ProductServerClient) interfaces.OrderService {
	return &OrderService{
		OrderRepository: orderRepository,
		ProductService:  productService,
	}
}

func (s *OrderService) CreateOrder(order *entity.Order) error {
	return s.OrderRepository.CreateOrder(order)
}

func (s *OrderService) DeleteOrder(id string) error {
	return s.OrderRepository.DeleteOrder(id)
}
