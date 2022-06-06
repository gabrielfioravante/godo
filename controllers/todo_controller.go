package controllers

import (
	. "godo/db"
	"godo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	Title       string `json:"title" binding:"required,min=2,max=50"`
	Description string `json:"description" binding:"required,min=2"`
	ListID      int    `json:"listID" binding:"required"`
}

func CreateTodo(c *gin.Context) {
	var input todo
	var list models.List

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if DB.First(&list, input.ListID).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "List not found"})
		return
	}

	todo := models.Todo{Title: input.Title, Description: input.Description}
	DB.Model(&list).Association("Todos").Append(&todo)

	c.JSON(http.StatusCreated, gin.H{"data": todo})
}

func DeleteTodo(c *gin.Context) {
	var todo models.List

	if DB.First(&todo, c.Param("id")).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	DB.Delete(&todo)

	c.JSON(http.StatusOK, gin.H{})
}

type updateTodo struct {
	Title       *string `json:"title" binding:"omitempty,min=2,max=50"`
	Description *string `json:"description" binding:"omitempty,min=2"`
}

func UpdateTodo(c *gin.Context) {
	var input updateTodo
	var todo models.Todo

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if DB.First(&todo, c.Param("id")).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "List not found"})
		return
	}

	if input.Title != nil {
		todo.Title = *input.Title
	}
	if input.Description != nil {
		todo.Description = *input.Description
	}

	DB.Save(&todo)

	c.JSON(http.StatusOK, gin.H{})
}
