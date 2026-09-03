package routes

import (
	"todo-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Todo API berjalan",
		})
	})

	r.GET("/tasks", handlers.GetTasks)
	r.GET("/tasks/:id", handlers.GetTaskByID)
	r.POST("/tasks", handlers.CreateTask)
	r.PUT("/tasks/:id", handlers.UpdateTask)
	r.DELETE("/tasks/:id", handlers.DeleteTask)
}
