package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Gin *gin.Engine
}

func NewHandler() *Handler {
	return &Handler{
		Gin: gin.Default(),
	}
}