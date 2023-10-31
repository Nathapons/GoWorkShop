package db

import (
	"main/models"

	_"gorm.io/driver/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func SetupDB() {
	dsn := "host=localhost user=admin password=12345678 dbname=golang port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Product{})
	database.AutoMigrate(&models.Transaction{})


	db = database
}
