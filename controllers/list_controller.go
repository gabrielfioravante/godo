package controllers

import (
	. "godo/db"
	"godo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateList(c *gin.Context) {
	var input models.ListAPIPost

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	list := models.List{Title: input.Title}
	DB.Create(&list)

	c.JSON(http.StatusCreated, gin.H{"data": list})
}

func GetLists(c *gin.Context) {
	var list models.List
	var lists []models.ListAPIGet

	DB.Model(&list).Find(&lists)

	c.JSON(http.StatusOK, gin.H{"data": lists})
}

func GetList(c *gin.Context) {
	var list models.List

	if DB.Preload("Todos").First(&list, c.Param("id")).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "List not found"})
		return
	}

	c.JSON(http.StatusFound, gin.H{"data": list})
}

func DeleteList(c *gin.Context) {
	var list models.List

	if DB.First(&list, c.Param("id")).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "List not found"})
		return
	}

	DB.Delete(&list)

	c.JSON(http.StatusOK, gin.H{})
}

func UpdateList(c *gin.Context) {
	var input models.ListAPIPost
	var list models.List

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if DB.First(&list, c.Param("id")).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "List not found"})
		return
	}

	DB.Model(&list).Update("title", input.Title)

	c.JSON(http.StatusOK, gin.H{})
}
