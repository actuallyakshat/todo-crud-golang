package routes

import (
	"github.com/actuallyakshat/todo-crud/controllers"
	"github.com/gin-gonic/gin"
)

func TodoRoutes(r *gin.Engine) {
	todoGroup := r.Group("/todos")
	{
		todoGroup.GET("/", controllers.GetTodos)
		todoGroup.POST("/", controllers.CreateTodo)
		todoGroup.PUT("/:id", controllers.UpdateTodo)
		todoGroup.DELETE("/:id", controllers.DeleteTodo)
		todoGroup.PATCH("/:id/complete", controllers.CompleteTodo)
	}
}
