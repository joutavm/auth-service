package controller

import (
	"auth-service/server/handler"
	userHandler "auth-service/user/handler"
)

func RegisterHandlers(handler *handler.Handler, userHandler *userHandler.UserHandler) {

	handler.Gin.POST("/user", userHandler.PostUser)
}
