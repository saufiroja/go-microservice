package config

import (
	"log"
	"os"
)

const (
	grpcDefaultHost = "localhost"
	grpcDefaultPort = "50053"
)

func initGRPC(conf *AppConfig) {
	host := os.Getenv("GRPC_HOST")
	port := os.Getenv("GRPC_PORT")

	switch {
	case host == "":
		log.Println("GRPC_HOST is not set, using default host", grpcDefaultHost)
		conf.GRPC.Host = grpcDefaultHost
	case port == "":
		log.Println("GRPC_PORT is not set, using default port", grpcDefaultPort)
		conf.GRPC.Port = grpcDefaultPort
	}

	conf.GRPC.Host = host
	conf.GRPC.Port = port
}
