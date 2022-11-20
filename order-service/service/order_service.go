package service

import (
	"github.com/saufiroja/order-service/entity"
	"github.com/saufiroja/order-service/interfaces"
)

type OrderService struct {
	OrderRepository interfaces.OrderRepository
}

func NewOrderService(orderRepository interfaces.OrderRepository) interfaces.OrderService {
	return &OrderService{
		OrderRepository: orderRepository,
	}
}

func (s *OrderService) CreateOrder(order *entity.Order) error {
	return s.OrderRepository.CreateOrder(order)
}

func (s *OrderService) DeleteOrder(id string) error {
	return s.OrderRepository.DeleteOrder(id)
}
