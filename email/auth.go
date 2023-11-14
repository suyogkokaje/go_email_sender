package email

import (
	"net/smtp"

	"go_email_sender/config"
)

func CreateAuth(configs *config.EmailConfig) smtp.Auth {
	return smtp.PlainAuth("", configs.SenderEmail, configs.SecretKey, configs.SmtpHost)
}
