package migrate

import (
	"auth-service/auth/model"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
)

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalln("Failed migration")
	}
}

func RegisterMigration(db *gorm.DB, config *viper.Viper) {
	if config.GetBool("db.migration") {
		migrate(db)
	}
}
