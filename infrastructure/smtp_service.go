package infrastructure

import (
	"strings"
	"sync"

	"gopkg.in/gomail.v2"
)

// SmtpService handles email sending using SMTP
type SmtpService struct {
	dialer      *gomail.Dialer
	mu          sync.Mutex
	htmlContent string
	senderName  string
	subject     string
}

// NewSmtpService creates a new instance of SmtpService
func NewSmtpService(smtpHost string, smtpPort int, senderEmail, password, htmlTemplate string, senderName string, subject string) *SmtpService {
	return &SmtpService{
		dialer:      gomail.NewDialer(smtpHost, smtpPort, senderEmail, password),
		htmlContent: htmlTemplate,
		senderName:  senderName,
		subject:     subject,
	}
}

// SetSenderName sets the sender's name dynamically
func (s *SmtpService) SetSenderName(name string) {
	s.senderName = name
}

// SetSubject sets the subject of the email dynamically
func (s *SmtpService) SetSubject(subject string) {
	s.subject = subject
}

// Send sends an email using the SMTP service
func (s *SmtpService) Send(email, username string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.senderName+"<"+s.dialer.Username+">")
	m.SetHeader("To", email)
	m.SetHeader("Subject", s.subject)

	// Replace placeholder with actual username
	if strings.Contains(s.htmlContent, "%name%") {
		body := strings.ReplaceAll(s.htmlContent, "%name%", username)
		m.SetBody("text/html", strings.TrimSpace(body))
	} else {
		m.SetBody("text/html", strings.TrimSpace(s.htmlContent))
	}

	// Locking the send operation to avoid race conditions
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.dialer.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
