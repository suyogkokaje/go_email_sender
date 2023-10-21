package email

import (
    "encoding/json"
    "io/ioutil"
)

type RecipientsConfig struct {
    Recipients []string `json:"recipients"`
}

func LoadRecipients(filename string) ([]string, error) {
    var config RecipientsConfig
    file, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    if err := json.Unmarshal(file, &config); err != nil {
        return nil, err
    }
    return config.Recipients, nil
}
