package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gitwooz/go-gin-app/services"
	"github.com/gitwooz/go-gin-app/utils"
)

type AuthController struct {
	AuthService services.AuthService
}

func (ac *AuthController) SignUp(c *gin.Context) {
	var input struct {
		PhoneNumber string `json:"phone_number" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid input")
		return
	}

	otp, err := ac.AuthService.GenerateOTP(input.PhoneNumber)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Error generating OTP")
		return
	}

	utils.SendResponse(c, http.StatusOK, "OTP sent", gin.H{"otp": otp})
}

func (ac *AuthController) SignIn(c *gin.Context) {
	var input struct {
		PhoneNumber string `json:"phone_number" binding:"required"`
		OTP         string `json:"otp" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid input")
		return
	}

	token, err := ac.AuthService.ValidateOTP(input.PhoneNumber, input.OTP)
	if err != nil {
		utils.SendError(c, http.StatusUnauthorized, "Invalid OTP")
		return
	}

	utils.SendResponse(c, http.StatusOK, "Sign in successful", gin.H{"token": token})
}
