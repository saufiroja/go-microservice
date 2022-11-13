package service

import (
	"errors"
	"fmt"
	"net/http"

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
	product, err := s.ProductService.FindOne(order.ProductID)
	if err != nil {
		return errors.New("[roduct not found")
	}
	fmt.Println(product)

	if int64(product.Data.Stock) < order.Quantity {
		return errors.New("product stock is not enough")
	}

	data := &entity.Order{
		Price:     int64(product.Data.Price),
		ProductID: product.Data.Id,
		UserID:    order.UserID,
	}

	err = s.OrderRepository.CreateOrder(data)
	if err != nil {
		return err
	}

	res, err := s.ProductService.DecreaseStock(order.ProductID, order.ID)
	if err != nil {
		return err
	}

	if res.Status == http.StatusConflict {
		return errors.New("Product stock is not enough")
	}

	return nil
}
