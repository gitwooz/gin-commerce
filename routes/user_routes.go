package routes

import (
    "github.com/gin-gonic/gin"
    "go-gin-app/controllers"
    "go-gin-app/middlewares"
)

func UserRoutes(router *gin.Engine) {
    userController := controllers.UserController{}

    userGroup := router.Group("/api/users")
    {
        userGroup.GET("/:id", userController.GetUserByID)
        userGroup.PUT("/:id", middlewares.JWTMiddleware(), userController.UpdateUser)
        userGroup.DELETE("/:id", middlewares.JWTMiddleware(), userController.DeleteUser)
    }
}