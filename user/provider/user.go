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

func Init(db *gorm.DB, encryptService *service.EncryptService) *UserProvider {
	return &UserProvider{
		db:      db,
		encrypt: encryptService}
}

func (userProvider UserProvider) CreateProduct(userRequest *model.User) (*model.User, error) {
	userRequest.Password = userProvider.encrypt.SHA256Encoder(userRequest.Password)
	tx := userProvider.db.Create(&userRequest)
	return userRequest, tx.Error
}
