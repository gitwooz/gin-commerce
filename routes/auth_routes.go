package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gitwooz/go-gin-app/controllers"
)

func AuthRoutes(router *gin.RouterGroup) {
	authController := controllers.AuthController{}

	router.POST("/signup", authController.SignUp)
	router.POST("/signin", authController.SignIn)
}
