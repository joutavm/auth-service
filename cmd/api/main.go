package main

import (
	"auth-service/auth"
	"auth-service/config"
	"auth-service/database"
	"auth-service/server"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		config.Module,
		database.Module,
		server.Module,
		auth.Module,
	).Run()
}
