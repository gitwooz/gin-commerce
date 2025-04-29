package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gitwooz/go-gin-app/config"
	"github.com/gitwooz/go-gin-app/routes"
	"github.com/gitwooz/go-gin-app/utils"

	files "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to MongoDB
	utils.ConnectMongoDB(cfg.DatabaseURL)

	// Initialize Gin router
	router := gin.Default()

	// Setup routes
	routes.UserRoutes(router)
	authGroup := router.Group("/api/auth")
	routes.AuthRoutes(authGroup)

	// Serve the swagger.yaml file
	router.StaticFile("/swagger.yaml", "./swagger.yaml")

	// Serve Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler, ginSwagger.URL("/swagger.yaml")))

	// Start the server
	router.Run(":" + cfg.Port)
}
