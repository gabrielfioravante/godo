package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title       string `gorm:"type:varchar(50)" json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	ListID      uint   `json:"listID" binding:"required"`
}
