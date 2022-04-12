package product

import (
	"go.uber.org/fx"
	"auth-service/product/controller"
	"auth-service/product/handler"
	"auth-service/product/migrate"
	"auth-service/product/provider"
)

var Module = fx.Options(
	migrate.Module,
	fx.Provide(
		provider.Init,
		handler.Init,
	),
	fx.Invoke(controller.RegisterHandlers),
)
