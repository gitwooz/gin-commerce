package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gitwooz/go-gin-app/controllers"
)

func UserRoutes(router *gin.Engine) {
	userController := controllers.UserController{}

	userGroup := router.Group("/api/users")
	{
		userGroup.POST("/", userController.CreateUser)      // Create a new user
		userGroup.GET("/:id", userController.GetUserByID)   // Get user by ID
		userGroup.PUT("/:id", userController.UpdateUser)    // Update user by ID
		userGroup.DELETE("/:id", userController.DeleteUser) // Delete user by ID
	}
}
