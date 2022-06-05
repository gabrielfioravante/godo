package server

import (
	"godo/controllers"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Setup(port int) {
	r := gin.Default()

	// List endpoints
	r.POST("/list", controllers.CreateList)
	r.GET("/list", controllers.GetLists)
	r.GET("/list/:id", controllers.GetList)
	r.DELETE("/list/:id", controllers.DeleteList)
	r.PUT("/list/:id", controllers.UpdateList)

	// Todo endpoints
	r.POST("/todo", controllers.CreateTodo)
	r.DELETE("/todo/:id", controllers.CreateTodo)
	r.PUT("/todo/:id", controllers.UpdateTodo)

	r.Run(":" + strconv.Itoa(port))
}
