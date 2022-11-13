package config

import (
	"log"
	"os"

	"github.com/saufiroja/order-service/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func initApp(conf *AppConfig) {
	env := os.Getenv("GO_ENV")

	switch cases.Lower(language.English).String(env) {
	case "development":
		conf.App.Env = "development"
		log.Println(utils.Color("green") + "App environment is set to development")
	case "production":
		conf.App.Env = "production"
		log.Println(utils.Color("green") + "App environment is set to production")
	default:
		conf.App.Env = "development"
		log.Println(utils.Color("blue")+"GO_ENV is not set, using default environment", "development")
	}
}
