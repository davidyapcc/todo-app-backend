package routes

import (
	"github.com/davidyapcc/todo-app-backend/handlers"
	"github.com/gin-gonic/gin"
)

func InitTodoRoutes(router *gin.Engine) {
	todoRouter := router.Group("/api/todos")
	{
		todoRouter.GET("/", handlers.GetAllTodos)
		todoRouter.GET("/:id", handlers.GetTodoByID)
		todoRouter.POST("/", handlers.CreateTodo)
		todoRouter.PUT("/:id", handlers.UpdateTodo)
		todoRouter.DELETE("/:id", handlers.DeleteTodo)
	}
}
