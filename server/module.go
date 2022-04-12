package server

import (
	"go.uber.org/fx"
	"auth-service/server/handler"
	"auth-service/server/hook"
)

var Module = fx.Options(
	fx.Provide(handler.NewHandler),
	fx.Invoke(hook.RegisterHooks),
)
