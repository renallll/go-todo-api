package routes

import (
	"todo-api/handlers"
	"todo-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Todo API berjalan",
		})
	})
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/tasks", handlers.GetTasks)
	protected.GET("/tasks/:id", handlers.GetTaskByID)
	protected.POST("/tasks", handlers.CreateTask)
	protected.PUT("/tasks/:id", handlers.UpdateTask)
	protected.DELETE("/tasks/:id", handlers.DeleteTask)
}
