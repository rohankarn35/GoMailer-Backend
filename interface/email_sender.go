package _interface

// EmailSender defines the contract for sending emails
type EmailSender interface {
	Send(email string, username string) error
}
