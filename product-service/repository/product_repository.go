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
