package main

import (
	"auth-service/config"
	"auth-service/database"
	"auth-service/product"
	"auth-service/server"
	"auth-service/user"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		config.Module,
		database.Module,
		server.Module,
		product.Module,
		user.Module,
	).Run()
}
