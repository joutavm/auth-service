package handler

import (
	"auth-service/auth/model"
	"auth-service/auth/provider"
	"auth-service/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserHandler struct {
	userProvider *provider.UserProvider
}

func (userHandler UserHandler) PostUser(context *gin.Context) {
	var userRequest model.User
	if err := context.ShouldBindJSON(&userRequest); err != nil {
		context.JSON(http.StatusBadRequest, response.BuildError(err.Error(), http.StatusBadRequest))
		return
	}
	user, err := userHandler.userProvider.CreateProduct(&userRequest)

	if err != nil {
		log.Println("Failed to save item", err)
		context.JSON(http.StatusInternalServerError, response.BuildInternalError())
		return
	}

	context.JSON(http.StatusCreated, user)
}

func InitUserHandler(userProvider *provider.UserProvider) *UserHandler {
	return &UserHandler{
		userProvider: userProvider,
	}
}
