package repository

import (
	"github.com/saufiroja/order-service/entity"
	"github.com/saufiroja/order-service/interfaces"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) interfaces.OrderRepository {
	return &OrderRepository{
		DB: db,
	}
}

func (r *OrderRepository) CreateOrder(order *entity.Order) error {
	res := r.DB.Model(&order).Create(&order)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *OrderRepository) DeleteOrder(id string) error {
	res := r.DB.Delete(&entity.Order{}, id)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
