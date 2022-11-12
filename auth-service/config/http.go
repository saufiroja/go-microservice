package config

import (
	"log"
	"os"
)

const httpDefaultPort = "8000"
const httpDefaultHost = "127.0.0.1"

func initHTTP(conf *AppConfig) {
	host := os.Getenv("HTTP_HOST")
	port := os.Getenv("HTTP_PORT")

	switch {
	case port == "":
		log.Println("HTTP_PORT is not set. Using default port " + httpDefaultPort)
		port = httpDefaultPort

	case host == "":
		log.Println("HTTP_HOST is not set. Using default host " + httpDefaultHost)
		host = httpDefaultHost
	}

	conf.HTTP.Host = host
	conf.HTTP.Port = port
}
