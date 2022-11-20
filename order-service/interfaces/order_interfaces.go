package interfaces

import "github.com/saufiroja/order-service/entity"

type OrderRepository interface {
	CreateOrder(order *entity.Order) error
	DeleteOrder(id string) error
}

type OrderService interface {
	CreateOrder(order *entity.Order) error
	DeleteOrder(id string) error
}
