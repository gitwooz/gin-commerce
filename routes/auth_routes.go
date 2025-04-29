package routes

import (
	"github.com/gin-gonic/gin"
	"your_project/controllers"
)

func AuthRoutes(router *gin.RouterGroup) {
	authController := controllers.NewAuthController()

	router.POST("/signup", authController.SignUp)
	router.POST("/signin", authController.SignIn)
}