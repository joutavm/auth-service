package main

import (
	"go.uber.org/fx"
	"auth-service/config"
	"auth-service/database"
	"auth-service/product"
	"auth-service/server"
)

func main() {
	fx.New(
		config.Module,
		database.Module,
		server.Module,
		product.Module,
	).Run()
}
