package config

import (
	"os"
)

func initPostgres(conf *AppConfig) {
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	conf.Postgres.DB_NAME = name
	conf.Postgres.DB_USER = user
	conf.Postgres.DB_PASS = pass
	conf.Postgres.DB_HOST = host
	conf.Postgres.DB_PORT = port
}
