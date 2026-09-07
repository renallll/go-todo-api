package main

import (
	"todo-api/database"
	"todo-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()

	router := gin.Default()
	routes.SetupRoutes(router)

	router.Run(":8080")
}
