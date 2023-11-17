package parser

import (
	"bytes"
	"io/ioutil"
	"log"
	"text/template"
)

// this struct data which will be passed in the HTML template
type TemplateData struct {
	Email  string
	Values map[string]string
}

// this takes data as an argument and pass it into the HTML template
// and parse the HTML into string format
func ParseHTML(filePath string, data TemplateData) string {

	htmlTemplate, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("error while reading the template: %s", err)
	}

	tmpl, err := template.New("emailTemplate").Parse(string(htmlTemplate))
	if err != nil {
		log.Printf("error while parsing the template: %s", err)
	}

	var buf = new(bytes.Buffer)
	err = tmpl.Execute(buf, data)
	if err != nil {
		log.Printf("error while converting html to string: %s", err)
	}

	htmlBody := buf.String()

	return htmlBody
}
