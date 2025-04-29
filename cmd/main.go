package main

import (
    "github.com/gin-gonic/gin"
    "go-gin-app/config"
    "go-gin-app/routes"
)

func main() {
    // Load configuration
    config.Load()

    // Initialize Gin router
    router := gin.Default()

    // Setup routes
    routes.SetupRoutes(router)

    // Start the server
    router.Run(":8080")
}