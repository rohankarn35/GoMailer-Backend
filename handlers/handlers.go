package handlers

import (
	"encoding/json"
	"fmt"
	"gomailer/domain"
	"gomailer/infrastructure"
	"gomailer/usecase"
	"net/http"
	"strings"
)

// SendEmailHandler handles the email sending HTTP request
func SendEmailHandler(w http.ResponseWriter, r *http.Request) {
	var req domain.EmailRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	fmt.Println(req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Create the SMTP email sender service from infrastructure
	smtpService := infrastructure.NewSmtpService(
		"smtp.gmail.com",
		587,
		req.SenderEmail,
		req.AppPassword,
		req.HTMLTemplate,
		req.SenderName,
		req.Subject,
	)

	// Set sender name and subject dynamically
	smtpService.SetSenderName(req.SenderName)
	smtpService.SetSubject(req.Subject)

	// Send emails using worker pool
	results := usecase.SendEmails(req.Users, smtpService, 10)

	if len(results) > 0 && strings.Contains(results[0].Message, "Username and Password not accepted") {
		w.WriteHeader(http.StatusUnauthorized)
		jsonResponse, err := json.Marshal(results)
		if err != nil {
			http.Error(w, "Failed to create response", http.StatusUnauthorized)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}

	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Failed to create response", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
