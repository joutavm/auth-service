package base

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
)
import "gorm.io/driver/postgres"

func Connect(config *viper.Viper) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.GetString("db.url")), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error connecting to database: ", err)
	}

	return db
}
