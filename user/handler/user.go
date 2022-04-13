package handler

import (
	"auth-service/user/model"
	"auth-service/user/provider"
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
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := userHandler.userProvider.CreateProduct(userRequest)

	if err != nil {
		log.Println("Failed to save item", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	context.JSON(http.StatusCreated, response)
}

func Init(userProvider *provider.UserProvider) *UserHandler {
	return &UserHandler{
		userProvider: userProvider,
	}
}
