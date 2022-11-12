package config

import (
	"log"
	"os"
)

const grpcDefaultPort = "8000"
const grpcDefaultHost = "127.0.0.1"

func initGrpc(conf *AppConfig) {
	host := os.Getenv("GRPC_HOST")
	port := os.Getenv("GRPC_PORT")

	switch {
	case port == "":
		log.Println("GRPC_PORT is not set. Using default port " + grpcDefaultPort)
		port = grpcDefaultPort

	case host == "":
		log.Println("GRPC_HOST is not set. Using default host " + grpcDefaultHost)
		host = grpcDefaultHost
	}

	conf.GRPC.Host = host
	conf.GRPC.Port = port
}
