package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gitwooz/go-gin-app/controllers"
	"github.com/gitwooz/go-gin-app/middlewares"
)

func UserRoutes(router *gin.Engine) {
	userController := controllers.UserController{}

	userGroup := router.Group("/api/users")
	{
		userGroup.GET("/:id", userController.GetUser)
		userGroup.PUT("/:id", middlewares.JWTMiddleware(), userController.UpdateUser)
	}
}
