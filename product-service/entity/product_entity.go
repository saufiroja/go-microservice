package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID               string           `json:"id"`
	Name             string           `json:"name"`
	Price            int32            `json:"price"`
	Stock            int32            `json:"stock"`
	StockDecreaseLog StockDecreaseLog `json:"stock_decrease_log" gorm:"foreignKey:ProductID"`
}

type StockDecreaseLog struct {
	ID        string `json:"id"`
	OrderID   string `json:"order_id"`
	ProductID string `json:"product_id"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New().String()
	return
}

func (s *StockDecreaseLog) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New().String()
	return
}
