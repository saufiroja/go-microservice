package config

import (
	"log"
	"os"

	"github.com/saufiroja/product-service/utils"
)

const (
	httpDefaultHost = "localhost"
	httpDefaultPort = "8080"
)

func initHttp(conf *AppConfig) {
	host := os.Getenv("HTTP_HOST")
	port := os.Getenv("HTTP_PORT")

	switch {
	case host == "":
		log.Println(utils.Color("yellow")+"HTTP_HOST is not set, using default host", httpDefaultHost)
		conf.HTTP.Host = httpDefaultHost
	case port == "":
		log.Println(utils.Color("yellow")+"HTTP_PORT is not set, using default port", httpDefaultPort)
		conf.HTTP.Port = httpDefaultPort
	}

	conf.HTTP.Host = host
	conf.HTTP.Port = port
}
