package config

import (
	"log"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	App struct {
		Env string
	}
	HTTP struct {
		Host string
		Port string
	}
	GRPC struct {
		Host string
		Port string
	}
	Postgres struct {
		DB_HOST string
		DB_PORT string
		DB_USER string
		DB_PASS string
		DB_NAME string
	}
}

var appConfig *AppConfig

func NewAppConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}

	if appConfig == nil {
		appConfig = &AppConfig{}

		initApp(appConfig)
		initHTTP(appConfig)
		initGrpc(appConfig)
		initPostgres(appConfig)
	}

	return appConfig
}
