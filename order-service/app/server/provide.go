package server

import (
	"github.com/saufiroja/order-service/app/client"
	"github.com/saufiroja/order-service/app/handler"
	"github.com/saufiroja/order-service/config"
	"github.com/saufiroja/order-service/database/postgres"
	"github.com/saufiroja/order-service/repository"
	"github.com/saufiroja/order-service/service"
)

type gRPCProvider struct {
	handlers struct {
		order handler.OrderHandler
	}
}

func (rpc *GrpcServer) Provide() gRPCProvider {
	provider := gRPCProvider{}

	conf := config.NewAppConfig()
	db := postgres.NewPostgres(conf)
	client := client.NewProductServerClient()

	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo, client)
	provider.handlers.order = *handler.NewOrderHandler(orderService, client)

	return provider
}
