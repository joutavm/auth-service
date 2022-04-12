package provider

import (
	"gorm.io/gorm"
	"auth-service/product/model"
	"auth-service/product/request"
)

type ProductProvider struct {
	db *gorm.DB
}

func (productProvider ProductProvider) CreateProduct(productRequest *request.ProductRequest) (*model.Product, error) {
	product := &model.Product{
		Name:  productRequest.Name,
		Price: productRequest.Price,
	}
	tx := productProvider.db.Create(&product)
	return product, tx.Error
}

func (productProvider ProductProvider) GetAllProduct() (*[]model.Product, error) {
	var products []model.Product
	tx := productProvider.db.Find(&products)
	return &products, tx.Error

}

func (productProvider ProductProvider) GetById(id uint64) (*model.Product, error) {
	var products model.Product
	tx := productProvider.db.First(&products, id)
	return &products, tx.Error

}

func (productProvider ProductProvider) Update(id uint, patch *request.ProductPatch) (*model.Product, int64, error) {
	productModel := model.Product{
		Model: gorm.Model{
			ID: id,
		},
		Name:  patch.Name,
		Price: patch.Price,
	}
	tx := productProvider.db.Save(&productModel)
	return &productModel, tx.RowsAffected, tx.Error
}

func Init(db *gorm.DB) *ProductProvider {
	return &ProductProvider{
		db: db,
	}
}
