package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	EmailConfirmation EmailConfirmationConfig
}

type EmailConfirmationConfig struct {
	Email    string
	Password string
	Address  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default config")
	}

	return &Config{
		EmailConfirmation: EmailConfirmationConfig{
			Email:    os.Getenv("CONFIRMATION_EMAIL"),
			Password: os.Getenv("CONFIRMATION_PASS"),
			Address:  os.Getenv("CONFIRMATION_ADDRESS"),
		},
	}
}
