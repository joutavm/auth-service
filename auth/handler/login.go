package handler

import (
	"auth-service/auth/provider"
	"auth-service/auth/request"
	"auth-service/auth/service"
	"auth-service/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginHandler struct {
	userProvider *provider.UserProvider
	encrypt      *service.EncryptService
	jwtService   *service.JwtService
}

func InitLoginHandler(
	userProvider *provider.UserProvider,
	encryptService *service.EncryptService,
	jwtService *service.JwtService) *LoginHandler {

	return &LoginHandler{
		userProvider: userProvider,
		encrypt:      encryptService,
		jwtService:   jwtService,
	}
}

func (h *LoginHandler) Login(c *gin.Context) {
	var login request.Login

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, response.BuildError(err.Error(), http.StatusBadRequest))
		return
	}

	user, err := h.userProvider.GetUserByEmail(login.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			response.BuildError(fmt.Sprintf("User with email %s not found", login.Email), http.StatusBadRequest))
		return
	}

	if user.Password != h.encrypt.SHA256Encoder(login.Password) {
		c.JSON(http.StatusUnauthorized, response.BuildError("Wrong password", http.StatusUnauthorized))
		return
	}

	token, err := h.jwtService.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.BuildInternalError())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}
