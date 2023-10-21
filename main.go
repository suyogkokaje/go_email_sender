package main

import (
	"log"
	"os"
	"go_email_sender/config"
	"go_email_sender/email"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatal("Failed to load configuration")
	}

	senderEmail := os.Getenv("SENDER_EMAIL")
	secretKey := os.Getenv("SECRET_KEY")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	recipients, err := email.LoadRecipients("recipients.json")
	if err != nil {
		log.Fatal("Failed to load recipients:", err)
	}

	auth := email.CreateAuth(senderEmail, secretKey, smtpHost)

	emailContent, err := email.LoadEmailContent("email_content.json")
	if err != nil {
		log.Fatal("Failed to load email content:", err)
	}

	for _, recipient := range recipients {
		if err := email.SendEmail(senderEmail, smtpHost, smtpPort, auth, recipient, emailContent); err != nil {
			log.Printf("Error sending email to %s: %v", recipient, err)
		} else {
			log.Printf("Email sent to %s", recipient)
		}
	}
}
