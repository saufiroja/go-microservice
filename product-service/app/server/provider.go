package server

import (
	"github.com/saufiroja/product-service/app/handler"
	"github.com/saufiroja/product-service/config"
	"github.com/saufiroja/product-service/database/postgres"
	"github.com/saufiroja/product-service/repository"
	"github.com/saufiroja/product-service/service"
)

type gRPCProvider struct {
	handlers struct {
		product handler.ProductHandler
	}
}

func (rpc *GrpcServer) provide() gRPCProvider {
	provider := gRPCProvider{}

	conf := config.NewAppConfig()
	db := postgres.NewPostgres(conf)

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	provider.handlers.product = *handler.NewProductHandler(productService)

	return provider
}
