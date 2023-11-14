package main

import (
	"log"
	config "go_email_sender/config"
	"go_email_sender/email"
)

func main() {

	configs := config.ExtractConfig()

	recipients, err := email.LoadRecipients("recipients.json")
	if err != nil {
		log.Fatal("Failed to load recipients:", err)
	}

	auth := email.CreateAuth(configs)

	emailContent, err := email.LoadEmailContent("email_content.json")
	if err != nil {
		log.Fatal("Failed to load email content:", err)
	}

	for _, recipient := range recipients {
		if err := email.SendEmail(configs.SenderEmail, configs.SmtpHost, configs.SmtpPort, auth, recipient, emailContent); err != nil {
			log.Printf("Error sending email to %s: %v", recipient, err)
		} else {
			log.Printf("Email sent to %s", recipient)
		}
	}
}
