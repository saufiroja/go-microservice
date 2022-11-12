package main

import (
	"github.com/saufiroja/auth-service/app/server"
	"github.com/saufiroja/auth-service/config"
)

func main() {
	config := config.NewAppConfig()

	app := server.NewGrpcServer(config.GRPC.Host, config.GRPC.Port)

	app.Start()
}
