package models

import "gorm.io/gorm"

type List struct {
	gorm.Model
	Title string `gorm:"type:varchar(50)" json:"title"`
	Todos []Todo `json:"todos"`
}

type APIList struct {
	gorm.Model
	Title string `json:"title"`
}
