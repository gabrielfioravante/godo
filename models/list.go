package models

import "gorm.io/gorm"

type List struct {
	gorm.Model
	Title string `gorm:"type:varchar(50)" json:"title"`
	Todos []Todo `json:"todos"`
}

type ListAPIGet struct {
	gorm.Model
	Title string `json:"title"`
}

type ListAPIPost struct {
	Title string `json:"title" binding:"required,min=2,max=50"`
}
