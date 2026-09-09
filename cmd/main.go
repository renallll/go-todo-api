// @title Todo API
// @version 1.0
// @description Todo API menggunakan Go Gin PostgreSQL JWT
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Masukkan JWT token
package main

import (
	"todo-api/database"
	_ "todo-api/docs"
	"todo-api/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	database.ConnectDatabase() // WAJIB sebelum router

	router := gin.Default()

	routes.SetupRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")

}
