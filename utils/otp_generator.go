package utils

import (
	"math/rand"
	"strconv"
	"time"
)

const otpLength = 6

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateOTP() string {
	otp := rand.Intn(1000000) // Generate a random number between 0 and 999999
	return strconv.Itoa(otp)[len(strconv.Itoa(otp))-otpLength:] // Ensure it's 6 digits
}

func ValidateOTP(inputOTP string, generatedOTP string) bool {
	return inputOTP == generatedOTP
}