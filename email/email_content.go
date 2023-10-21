package email

import (
    "encoding/json"
    "io/ioutil"
)

type EmailContent struct {
    Subject string `json:"subject"`
    Body    string `json:"body"`
}

func LoadEmailContent(filename string) (EmailContent, error) {
    var emailContent EmailContent
    file, err := ioutil.ReadFile(filename)
    if err != nil {
        return EmailContent{}, err
    }
    if err := json.Unmarshal(file, &emailContent); err != nil {
        return EmailContent{}, err
    }
    return emailContent, nil
}
