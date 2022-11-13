package service

import (
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
