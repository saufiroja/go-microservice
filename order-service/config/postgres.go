package config

import "os"

func initPostgres(conf *AppConfig) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	conf.Postgres.DB_HOST = dbHost
	conf.Postgres.DB_PORT = dbPort
	conf.Postgres.DB_USER = dbUser
	conf.Postgres.DB_PASS = dbPass
	conf.Postgres.DB_NAME = dbName
}
