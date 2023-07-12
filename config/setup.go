package config

import (
	"simple-go-restapi/pkg/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/simple-go-restapi"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&models.Product{})

	DB = database
	return DB
}
