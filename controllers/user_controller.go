package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-gin-app/models"
	"go-gin-app/services"
)

type UserController struct {
	UserService services.UserService
}

func (uc *UserController) GetUser(c *gin.Context) {
	userID := c.Param("id")
	user, err := uc.UserService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	userID := c.Param("id")
	if err := uc.UserService.UpdateUser(userID, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}