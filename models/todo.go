package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title       string `gorm:"type:varchar(50)" json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Done        bool   `json:"done"`
	ListID      uint   `json:"listID" binding:"required"`
}

type TodoAPIPost struct {
	Title       string `json:"title" binding:"required,min=2,max=50"`
	Description string `json:"description" binding:"required,min=2"`
	ListID      int    `json:"listID" binding:"required"`
}

type TodoAPIUpdate struct {
	Title       *string `json:"title" binding:"omitempty,min=2,max=50"`
	Description *string `json:"description" binding:"omitempty,min=2"`
	Done        *bool   `json:"done" binding:"omitempty"`
}

