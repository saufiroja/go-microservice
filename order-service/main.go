package main

import (
	"github.com/saufiroja/order-service/app/server"
	"github.com/saufiroja/order-service/config"
)

func main() {
	config := config.NewAppConfig()

	app := server.NewGrpcServer(config.HTTP.Host, config.HTTP.Port)

	app.Start()
}
