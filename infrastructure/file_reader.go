package infrastructure

import (
	"encoding/json"
	"gomailer/domain"
	"io/ioutil"
	"log"
)

// LoadUsers reads the user list from a JSON file
func LoadUsers(filePath string) ([]domain.User, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading JSON file: %v", err)
		return nil, err
	}

	var jsonData struct {
		Users []domain.User `json:"users"`
	}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		return nil, err
	}

	return jsonData.Users, nil
}

// LoadTemplate reads the HTML template file
func LoadTemplate(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading HTML file: %v", err)
		return "", err
	}

	return string(content), nil
}
