package config

import "os"

func initApp(conf *AppConfig) {
	conf.App.Env = os.Getenv("GO_ENV")
}
