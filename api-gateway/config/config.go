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
	AuthService struct {
		URL string
	}
	ProductService struct {
		URL string
	}
	OrderService struct {
		URL string
	}
}

var app *AppConfig

func NewAppConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if app == nil {
		app = &AppConfig{}

		initApp(app)
		initAuthService(app)
		initProductService(app)
		initOrderService(app)
	}

	return app
}
