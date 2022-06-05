package db

import (
	"godo/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func Create() {
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	Instance = db
}

func Migrate() {
	Instance.AutoMigrate(&models.List{})
	Instance.AutoMigrate(&models.Todo{})
}
