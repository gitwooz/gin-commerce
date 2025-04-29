package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gitwooz/go-gin-app/models"
	"github.com/gitwooz/go-gin-app/services"
	"github.com/gitwooz/go-gin-app/utils"
)

type UserController struct {
	UserService services.UserService
}

func (uc *UserController) GetUser(c *gin.Context) {
	userID := c.Param("id")
	user, err := uc.UserService.GetUser(userID)
	if err != nil {
		utils.SendError(c, http.StatusNotFound, "User not found")
		return
	}
	utils.SendResponse(c, http.StatusOK, "User retrieved successfully", user)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid input")
		return
	}
	userID := c.Param("id")
	if err := uc.UserService.UpdateUser(userID, user); err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to update user")
		return
	}
	utils.SendResponse(c, http.StatusOK, "User updated successfully", nil)
}
