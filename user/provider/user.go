package provider

import (
	"auth-service/user/model"
	"auth-service/user/service"
	"gorm.io/gorm"
)

type UserProvider struct {
	encrypt *service.EncryptService
	db      *gorm.DB
}

func (userProvider UserProvider) CreateProduct(userRequest model.User) (*model.User, error) {
	userRequest.Password = userProvider.encrypt.SHA256Encoder(userRequest.Password)
	tx := userProvider.db.Save(&userRequest)
	return &userRequest, tx.Error
}

func Init(db *gorm.DB) *UserProvider {
	return &UserProvider{db: db}
}
