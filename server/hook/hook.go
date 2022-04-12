package hook

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"log"
	"auth-service/server/handler"
)

func RegisterHooks(lifecycle fx.Lifecycle, h *handler.Handler, config *viper.Viper) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			port := fmt.Sprintf(":%s", config.GetString("server.port"))
			go h.Gin.Run(port)
			return nil
		},
		OnStop:  func(ctx context.Context) error {
			log.Println("Stopping server")
			return nil },
	})
}

