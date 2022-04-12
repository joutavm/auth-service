package controller

import (
	productHandler "auth-service/product/handler"
	"auth-service/server/handler"
)

func RegisterHandlers(handler *handler.Handler, productHandler *productHandler.ProductHandler) {

	handler.Gin.GET("/product/:id", productHandler.GetProduct)
	handler.Gin.PUT("/product/:id", productHandler.PutProduct)
	handler.Gin.GET("/product", productHandler.GetProducts)
	handler.Gin.POST("/product", productHandler.PostProduct)
}
