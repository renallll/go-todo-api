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
	"os"

	"todo-api/database"
	_ "todo-api/docs"
	"todo-api/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	database.ConnectDatabase()

	// Otomatis Release Mode jika GIN_MODE=release
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// Hindari warning trusted proxies
	router.SetTrustedProxies(nil)

	routes.SetupRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")

}
