package config

import "os"

func initAuthService(conf *AppConfig) {
	conf.AuthService.URL = os.Getenv("AUTH_SERVICE_URL")
}

func initProductService(conf *AppConfig) {
	conf.ProductService.URL = os.Getenv("PRODUCT_SERVICE_URL")
}

func initOrderService(conf *AppConfig) {
	conf.OrderService.URL = os.Getenv("ORDER_SERVICE_URL")
}
