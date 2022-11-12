package server

import (
	"github.com/saufiroja/auth-service/app/handler"
	"github.com/saufiroja/auth-service/config"
	"github.com/saufiroja/auth-service/database/postgres"
	"github.com/saufiroja/auth-service/repository"
	"github.com/saufiroja/auth-service/service"
)

type gRPCProvider struct {
	handlers struct {
		user handler.UserHandler
	}
}

func (rpc *GrpcServer) provide() gRPCProvider {
	provider := gRPCProvider{}

	conf := config.NewAppConfig()
	db := postgres.NewPostgres(conf)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	provider.handlers.user = *handler.NewUserHandler(userService)

	return provider
}
