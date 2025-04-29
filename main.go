package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gitwooz/go-gin-app/config"
	"github.com/gitwooz/go-gin-app/routes"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize Gin router
	router := gin.Default()

	// Setup routes
	routes.UserRoutes(router)
	authGroup := router.Group("/api/auth")
	routes.AuthRoutes(authGroup)

	// Start the server
	router.Run(":" + cfg.Port)
}
