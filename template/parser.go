package template

import (
	"bytes"
	"io/ioutil"
	"log"
	"text/template"
)

func ParseHTML() string{

	//this struct represents the data which you want to pass in your template.html file
	type TemplateData struct{
		Name string
	}

	//pass the actual values of the data here
	data := TemplateData{
		Name : "Suyog",
	}

	htmlTemplate, err := ioutil.ReadFile("/home/suyog/Coding/web/go_email_sender/template/template.html")
	if err != nil {
		log.Fatal("error while reading the template")
	}

	tmpl,err := template.New("emailTemplate").Parse(string(htmlTemplate))
	if err!=nil{
		log.Fatal("error while parsing the template")
	}

	var buf = new(bytes.Buffer)
	err = tmpl.Execute(buf,data)
	if err!=nil{
		log.Fatal("error while converting html to string")
	}

	htmlBody := buf.String()

	return htmlBody
}