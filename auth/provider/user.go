package provider

import (
	"auth-service/auth/model"
	"auth-service/auth/service"
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

func (userProvider UserProvider) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	tx := userProvider.db.Where("email = ?", email).First(&user)
	return &user, tx.Error
}
