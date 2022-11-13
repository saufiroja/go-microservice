package interfaces

import "github.com/saufiroja/product-service/entity"

type ProductRepository interface {
	CreateProduct(product *entity.Product) error
	FindOne(id string) (*entity.Product, error)
}

type ProductService interface {
	CreateProduct(product *entity.Product) error
	FindOne(id string) (*entity.Product, error)
}
