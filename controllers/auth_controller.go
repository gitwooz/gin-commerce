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
		c.JSON(http.StatusBadRequest, utils.ResponseUtil("Invalid input", nil))
		return
	}

	otp, err := ac.AuthService.GenerateOTP(input.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseUtil("Error generating OTP", nil))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseUtil("OTP sent", gin.H{"otp": otp}))
}

func (ac *AuthController) SignIn(c *gin.Context) {
	var input struct {
		PhoneNumber string `json:"phone_number" binding:"required"`
		OTP         string `json:"otp" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseUtil("Invalid input", nil))
		return
	}

	token, err := ac.AuthService.ValidateOTP(input.PhoneNumber, input.OTP)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.ResponseUtil("Invalid OTP", nil))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseUtil("Sign in successful", gin.H{"token": token}))
}
