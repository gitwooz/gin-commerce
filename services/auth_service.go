package services

import (
    "errors"
    "time"

    "github.com/golang-jwt/jwt/v4"
)

type AuthService struct{}

var jwtSecret = []byte("your-secret-key")

func (as *AuthService) GenerateOTP(phoneNumber string) (string, error) {
    // Simulate OTP generation
    return "123456", nil
}

func (as *AuthService) ValidateOTP(phoneNumber, otp string) (string, error) {
    // Simulate OTP validation
    if otp != "123456" {
        return "", errors.New("invalid OTP")
    }

    // Generate JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "phone_number": phoneNumber,
        "exp":          time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}