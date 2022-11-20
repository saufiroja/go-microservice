package main

import (
	"log"

	"github.com/saufiroja/api-gateway/app/server"
)

func main() {
	app := server.NewFiberServer()

	err := app.Start(":3000")

	if err != nil {
		log.Panic(err)
	}
}
