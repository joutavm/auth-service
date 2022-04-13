package user

import (
	"auth-service/user/controller"
	"auth-service/user/handler"
	"auth-service/user/migrate"
	"auth-service/user/provider"
	"auth-service/user/service"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		provider.Init,
		service.InitEncryptService,
		handler.Init),
	fx.Invoke(migrate.RegisterMigration),
	fx.Invoke(controller.RegisterHandlers),
)
