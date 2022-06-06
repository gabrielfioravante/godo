package db

import (
	"godo/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Create() {
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	DB = db
}

func Migrate() {
	DB.AutoMigrate(&models.List{})
	DB.AutoMigrate(&models.Todo{})
}
