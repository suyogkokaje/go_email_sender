package email

import (
	"go_email_sender/config"
	"net/smtp"
)

func SendEmail(configs *config.EmailConfig, auth smtp.Auth, recipient string, emailContent EmailContent) error {
	to := []string{recipient}

	msg := []byte("To: " + recipient + "\r\n" +
		"Subject: " + emailContent.Subject + "\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n" +
		emailContent.Body)

	err := smtp.SendMail(configs.SmtpHost+":"+configs.SmtpPort, auth, configs.SenderEmail, to, msg)
	return err
}
