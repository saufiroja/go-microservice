package config

import (
	"log"
	"os"
)

const (
	httpDefaultHost = "localhost"
	httpDefaultPort = "8000"
)

func initHTTP(conf *AppConfig) {
	host := os.Getenv("HTTP_HOST")
	port := os.Getenv("HTTP_PORT")

	switch {
	case host == "":
		log.Println("HTTP_HOST is not set, using default host", httpDefaultHost)
		conf.HTTP.Host = httpDefaultHost
	case port == "":
		log.Println("HTTP_PORT is not set, using default port", httpDefaultPort)
		conf.HTTP.Port = httpDefaultPort
	}

	conf.HTTP.Host = host
	conf.HTTP.Port = port
}
