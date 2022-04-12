package base

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
)
import "gorm.io/driver/postgres"


func Connect(config *viper.Viper) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.GetString("db.host"),
		config.GetString("db.user"),
		config.GetString("db.pass"),
		config.GetString("db.name"),
		config.GetString("db.port"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error connecting to database: ", err)
	}

	return db
}
