package controller

import (
	handlers "auth-service/auth/handler"
	baseHandler "auth-service/server/handler"
)

func RegisterHandlers(handler *baseHandler.Handler,
	userHandler *handlers.UserHandler,
	loginHandler *handlers.LoginHandler) {

	handler.Gin.POST("/user", userHandler.PostUser)
	handler.Gin.POST("/login", loginHandler.Login)
}
