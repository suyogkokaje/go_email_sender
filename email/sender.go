package email

import (
    "net/smtp"
)



func SendEmail(senderEmail, smtpHost, smtpPort string, auth smtp.Auth, recipient string, emailContent EmailContent) error {
    to := []string{recipient}

    msg := []byte("To: " + recipient + "\r\n" +
        "Subject: " + emailContent.Subject + "\r\n" +
        "\r\n" +
        emailContent.Body)

    err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, to, msg)
    return err
}
