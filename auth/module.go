package auth

import (
	"auth-service/auth/controller"
	"auth-service/auth/handler"
	"auth-service/auth/migrate"
	"auth-service/auth/provider"
	"auth-service/auth/service"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		provider.Init,
		service.InitEncryptService,
		handler.InitUserHandler,
		service.NewJwtService,
		handler.InitLoginHandler),
	fx.Invoke(migrate.RegisterMigration),
	fx.Invoke(controller.RegisterHandlers),
)
