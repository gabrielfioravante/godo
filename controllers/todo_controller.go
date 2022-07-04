package controllers

import (
	. "godo/db"
	"godo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func CreateTodo(c *gin.Context) {
	var input models.TodoAPIPost
	var list models.List

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if DB.First(&list, input.ListID).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "List not found"})
		return
	}

	todo := models.Todo{Title: input.Title, Description: input.Description, Done: false}
	DB.Model(&list).Association("Todos").Append(&todo)

	c.JSON(http.StatusCreated, gin.H{"data": todo})
}

func DeleteTodo(c *gin.Context) {
	var todo models.Todo

	if DB.First(&todo, c.Param("id")).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	DB.Delete(&todo)

	c.JSON(http.StatusOK, gin.H{})
}

func UpdateTodo(c *gin.Context) {
	var input models.TodoAPIUpdate
	var todo models.Todo

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if DB.First(&todo, c.Param("id")).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	models.UpdateFields(input, &todo)

	DB.Save(&todo)

	c.JSON(http.StatusOK, gin.H{})
}
