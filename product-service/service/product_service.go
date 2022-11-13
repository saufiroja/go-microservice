package service

import (
	"errors"

	"github.com/saufiroja/product-service/entity"
	"github.com/saufiroja/product-service/interfaces"
)

type ProductService struct {
	ProductRepository interfaces.ProductRepository
}

func NewProductService(productRepository interfaces.ProductRepository) interfaces.ProductService {
	return &ProductService{ProductRepository: productRepository}
}

func (p *ProductService) CreateProduct(product *entity.Product) error {
	return p.ProductRepository.CreateProduct(product)
}

func (p *ProductService) FindOne(id string) (*entity.Product, error) {
	return p.ProductRepository.FindOne(id)
}

func (p *ProductService) DecreaseStock(id, orderID string) error {
	product, err := p.ProductRepository.FindOne(id)
	if err != nil {
		return err
	}

	if product.Stock <= 0 {
		return errors.New("stock is empty")
	}

	log, err := p.ProductRepository.FindStockLog(orderID)
	if err != nil {
		return errors.New("stock already decreased")
	}

	product.Stock = product.Stock - 1

	err = p.ProductRepository.SaveProduct(product)
	if err != nil {
		return err
	}

	log.ProductID = product.ID
	log.OrderID = orderID

	err = p.ProductRepository.CreateStockLog(log)
	if err != nil {
		return err
	}

	return nil
}
