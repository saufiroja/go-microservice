package interfaces

import "github.com/saufiroja/product-service/entity"

type ProductRepository interface {
	CreateProduct(product *entity.Product) error
	FindOne(id string) (*entity.Product, error)
	FindStockLog(orderId string) (*entity.StockDecreaseLog, error)
	SaveProduct(product *entity.Product) error
	CreateStockLog(log *entity.StockDecreaseLog) error
}

type ProductService interface {
	CreateProduct(product *entity.Product) error
	FindOne(id string) (*entity.Product, error)
	DecreaseStock(id, orderID string) error
}
