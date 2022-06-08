package server

import (
	"godo/controllers"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Setup(mode string, port int) {
	gin.SetMode(mode)
	r := gin.Default()

	list := r.Group("/list")
	{
		list.POST("", controllers.CreateList)
		list.GET("", controllers.GetLists)
		list.GET("/:id", controllers.GetList)
		list.DELETE("/:id", controllers.DeleteList)
		list.PUT("/:id", controllers.UpdateList)
	}

	todo := r.Group("/todo")
	{
		todo.POST("", controllers.CreateTodo)
		todo.DELETE("/:id", controllers.DeleteTodo)
		todo.PUT("/:id", controllers.UpdateTodo)
	}

	r.Run(":" + strconv.Itoa(port))
}
