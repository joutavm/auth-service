package database

import (
	"go.uber.org/fx"
	"auth-service/database/base"
)

var Module = fx.Options(fx.Provide(base.Connect))
