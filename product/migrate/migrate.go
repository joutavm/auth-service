package migrate

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"log"
	"auth-service/product/model"
)

var Module = fx.Options(
	fx.Invoke(RegisterMigration))

func migrate(db *gorm.DB) 	{
	err := db.AutoMigrate(&model.Product{})
	if err != nil {
		log.Fatalln("Failed migration")
	}
}

func RegisterMigration(db *gorm.DB, config *viper.Viper) 	{
	if config.GetBool("db.migration") {
		migrate(db)
	}
}
