package main

import (
	"go_email_sender/config"
	"go_email_sender/email"
	"go_email_sender/parser"
	"log"
)

func main() {

	configs := config.ExtractConfig()

	auth := email.CreateAuth(configs)

	csvPath := "/home/suyog/Coding/web/go_email_sender/data/emaildata.csv"

	templatePath := "/home/suyog/Coding/web/go_email_sender/template/template.html"

	emailContent, err := email.LoadEmailContent("email_content.json")
	if err != nil {
		log.Fatal("Failed to load email content:", err)
	}

	headers, records, err := parser.ParseCSV(csvPath)

	for _, record := range records {
		mp := make(map[string]string)

		for i, val := range record {
			mp[headers[i]] = val
		}

		data := parser.TemplateData{
			Email:  mp["email"],
			Values: mp,
		}

		err := email.SendEmail(configs, auth, templatePath, data, emailContent)
		if err != nil {
			log.Fatal("error while sending an email")
		} else {
			log.Println("email delivered successfully")
		}
	}
}
