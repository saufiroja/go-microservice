package main

import (
	"github.com/saufiroja/product-service/app/server"
	"github.com/saufiroja/product-service/config"
)

func main() {
	config := config.NewAppConfig()

	app := server.NewGrpcServer(config.GRPC.Host, config.GRPC.Port)
	app.Start()
}
