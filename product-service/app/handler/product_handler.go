package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/saufiroja/product-service/entity"
	"github.com/saufiroja/product-service/interfaces"
	"github.com/saufiroja/product-service/rpc/pb/product"
)

type ProductHandler struct {
	product.UnimplementedProductServiceServer
	product interfaces.ProductService
}

func NewProductHandler(product interfaces.ProductService) *ProductHandler {
	return &ProductHandler{product: product}
}

func (p *ProductHandler) CreateProduct(ctx context.Context, req *product.CreateProductRequest) (*product.CreateProductResponse, error) {
	data := &entity.Product{
		Name:  req.Name,
		Price: req.Price,
		Stock: req.Stock,
	}
	err := p.product.CreateProduct(data)
	if err != nil {
		return nil, err
	}

	return &product.CreateProductResponse{
		Id:     data.ID,
		Status: fiber.StatusCreated,
	}, nil
}

func (p *ProductHandler) FindOne(ctx context.Context, req *product.FindOneRequest) (*product.FindOneResponse, error) {
	data, err := p.product.FindOne(req.Id)
	if err != nil {
		return nil, err
	}

	result := &product.FindOneData{
		Id:    data.ID,
		Name:  data.Name,
		Price: data.Price,
		Stock: data.Stock,
	}

	return &product.FindOneResponse{
		Data:   result,
		Status: fiber.StatusOK,
	}, nil
}
