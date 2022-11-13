package repository

import (
	"github.com/saufiroja/product-service/entity"
	"github.com/saufiroja/product-service/interfaces"
	"gorm.io/gorm"
)

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) interfaces.ProductRepository {
	return &productRepository{DB: db}
}

func (p *productRepository) CreateProduct(product *entity.Product) error {
	res := p.DB.Model(&product).Create(&product)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (p *productRepository) FindOne(id string) (*entity.Product, error) {
	var product entity.Product
	res := p.DB.Model(&product).Where("id = ?", id).First(&product)
	if res.Error != nil {
		return nil, res.Error
	}

	return &product, nil
}

func (p *productRepository) FindStockLog(orderId string) (*entity.StockDecreaseLog, error) {
	var log entity.StockDecreaseLog

	result := p.DB.Model(&log).Where("order_id = ?", orderId).First(&log)
	if result.Error == nil {
		return nil, result.Error
	}

	return &log, nil
}

func (p *productRepository) SaveProduct(product *entity.Product) error {
	res := p.DB.Model(&product).Save(&product)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (p *productRepository) CreateStockLog(log *entity.StockDecreaseLog) error {
	err := p.DB.Model(&log).Create(&log).Error
	if err != nil {
		return err
	}

	return nil
}
