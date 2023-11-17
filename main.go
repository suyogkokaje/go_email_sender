package main

import (
	"encoding/csv"
	"go_email_sender/config"
	"go_email_sender/email"
	"go_email_sender/template"
	"log"
	"os"
)



func main() {

	configs := config.ExtractConfig()

	// recipients, err := email.LoadRecipients("recipients.json")
	// if err != nil {
	// 	log.Fatal("Failed to load recipients:", err)
	// }

	auth := email.CreateAuth(configs)

	emailContent, err := email.LoadEmailContent("email_content.json")
	if err != nil {
		log.Fatal("Failed to load email content:", err)
	}

	// for _, recipient := range recipients {
	// 	if err := email.SendEmail(configs, auth, recipient, emailContent); err != nil {
	// 		log.Printf("Error sending email to %s: %v", recipient, err)
	// 	} else {
	// 		log.Printf("Email sent to %s", recipient)
	// 	}
	// }
	
	file,err := os.Open("/home/suyog/Coding/web/go_email_sender/data/emaildata.csv")
	if err!=nil{
		log.Fatal("error in opening the csv file")
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records,err := reader.ReadAll()
	if err!=nil{
		log.Fatal("error while reading all the records from the csv file")
	}

	headers := records[0]

	for _,record := range records[1:]{
		mp := make(map[string]string)

		for i,val:= range record{
			mp[headers[i]] = val
		}

		data := template.TemplateData{
			Email: mp["email"],
			Values: mp,
		}

		err := email.SendEmail(configs,auth,data,emailContent)
		if err!=nil{
			log.Fatal("error while sending an email")
		}
	}


}
