package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	GoogleClientId     string
	GoogleClientSecret string
	GoogleRedirectURL  string
	JWTSecretKey       []byte
)

func LoadConfig() {
	// Load the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	GoogleClientId = os.Getenv("CLIENT_ID")
	GoogleClientSecret = os.Getenv("CLIENT_SECRET")
	GoogleRedirectURL = os.Getenv("REDIRECT_URL")
	JWTSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

	if GoogleClientId == "" || GoogleClientSecret == "" || GoogleRedirectURL == "" || len(JWTSecretKey) == 0 {
		log.Fatalf("Missing one or more required environment variables")
	}
}
