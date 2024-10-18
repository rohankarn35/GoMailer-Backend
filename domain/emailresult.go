package domain

// EmailResult represents the result of sending an email
type EmailResult struct {
	Email   string `json:"email"`
	Status  string `json:"status"` // "sent" or "failed"
	Message string `json:"message"`
}
