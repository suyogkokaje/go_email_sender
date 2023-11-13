package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// configs required to setup the smtp server
type EmailConfig struct {
	SenderEmail string
	SecretKey   string
	SmtpHost    string
	SmtpPort    string
}

// function to check whether the environment variables are there or not
func CheckConfig() {
	godotenv.Load()
	SenderEmail, _ := os.LookupEnv("SENDER_EMAIL")
	SecretKey, _ := os.LookupEnv("SECRET_KEY")
	SmtpHost, _ := os.LookupEnv("SMTP_HOST")
	SmtpPort, _ := os.LookupEnv("SMTP_PORT")

	if len(SenderEmail) == 0 || len(SecretKey) == 0 || len(SmtpHost) == 0 || len(SmtpPort) == 0 {
		log.Fatal("Values of required environment variables are not provided!")
	}
}

// function to extract the values of the environment variable from .env file
func ExtractConfig() *EmailConfig {
	CheckConfig()

	configs := EmailConfig{
		SenderEmail: os.Getenv("SENDER_EMAIL"),
		SecretKey:   os.Getenv("SECRET_KEY"),
		SmtpHost:    os.Getenv("SMTP_HOST"),
		SmtpPort:    os.Getenv("SMTP_PORT"),
	}

	return &configs
}
