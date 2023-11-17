package email

import (
	"go_email_sender/config"
	"go_email_sender/template"
	"net/smtp"
)

func SendEmail(configs *config.EmailConfig, auth smtp.Auth, data template.TemplateData, emailContent EmailContent) error {

	//load the template
	emailBody := template.ParseHTML(data)

	to := []string{data.Email}

	msg := []byte("To: " + data.Email + "\r\n" +
		"Subject: " + emailContent.Subject + "\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n" +
		emailBody)

	err := smtp.SendMail(configs.SmtpHost+":"+configs.SmtpPort, auth, configs.SenderEmail, to, msg)
	return err
}
