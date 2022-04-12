package handler

import (
	"auth-service/product/provider"
	"auth-service/product/request"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	productProvider *provider.ProductProvider
}

func (productHandler ProductHandler) GetProduct(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Id must be a number"})
		return
	}
	product, err := productHandler.productProvider.GetById(id)
	if err != nil {
		log.Println("Failed to get item", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	context.JSON(http.StatusOK, product)
}

func (productHandler ProductHandler) GetProducts(context *gin.Context) {

	users, err := productHandler.productProvider.GetAllProduct()
	if err != nil {
		log.Println("Failed to get items", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	context.JSON(http.StatusOK, users)

}

func (productHandler ProductHandler) PostProduct(context *gin.Context) {
	var productRequest request.ProductRequest
	if err := context.ShouldBindJSON(&productRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := productHandler.productProvider.CreateProduct(&productRequest)

	if err != nil {
		log.Println("Failed to save item", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	context.JSON(http.StatusCreated, response)

}

func (productHandler ProductHandler) PutProduct(context *gin.Context) {
	var productPatch request.ProductPatch

	var err error
	var id uint64
	if err := context.ShouldBindJSON(&productPatch); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err = strconv.ParseUint(context.Param("id"), 10, 32)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Id must be a number"})
		return
	}

	response, rowsAffected, err := productHandler.productProvider.Update(uint(id), &productPatch)

	if err != nil {
		log.Println("Failed to update item", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	if rowsAffected == 0 {
		message := fmt.Sprintf("No item found with id %d", id)
		log.Println(message)
		context.JSON(http.StatusNotFound, gin.H{"error": message})
		return
	}
	context.JSON(http.StatusOK, response)
}

func Init(productProvider *provider.ProductProvider) *ProductHandler {
	return &ProductHandler{productProvider: productProvider}

}
