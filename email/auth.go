package email

import "net/smtp"

func CreateAuth(senderEmail, secretKey, smtpHost string) smtp.Auth {
	return smtp.PlainAuth("", senderEmail, secretKey, smtpHost)
}
