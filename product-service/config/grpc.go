package config

import (
	"log"
	"os"

	"github.com/saufiroja/product-service/utils"
)

const (
	grpcDefaultHost = "localhost"
	grpcDefaultPort = "50052"
)

func initGrpc(conf *AppConfig) {
	host := os.Getenv("GRPC_HOST")
	port := os.Getenv("GRPC_PORT")

	switch {
	case host == "":
		log.Println(utils.Color("yellow")+"GRPC_HOST is not set, using default host", grpcDefaultHost)
		conf.GRPC.Host = grpcDefaultHost
	case port == "":
		log.Println(utils.Color("yellow")+"GRPC_PORT is not set, using default port", grpcDefaultPort)
		conf.GRPC.Port = grpcDefaultPort
	}

	conf.GRPC.Host = host
	conf.GRPC.Port = port
}
