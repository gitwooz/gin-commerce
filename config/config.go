package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	JWTSecret    string
	DatabaseURL  string
	DatabaseName string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using default values")
	}

	return &Config{
		Port:         getEnv("PORT", "8080"),
		JWTSecret:    getEnv("JWT_SECRET", "your_jwt_secret"),
		DatabaseURL:  getEnv("DATABASE_URL", "mongodb://localhost:27017"),
		DatabaseName: getEnv("DATABASE_NAME", "gin-commerce"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
