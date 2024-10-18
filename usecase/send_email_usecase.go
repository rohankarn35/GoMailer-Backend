package usecase

import (
	"fmt"
	"gomailer/domain"
	"strings"
	"sync"
)

// EmailSender is an interface for sending emails
type EmailSender interface {
	Send(email, username string) error
}

// SendEmails handles the process of sending emails to users and returns success/failure results
func SendEmails(users []domain.User, sender EmailSender, maxWorkers int) []domain.EmailResult {
	user := users[0]
	err := sender.Send(user.Email, user.Name)
	results := make([]domain.EmailResult, 0, len(users))

	if err != nil {
		if strings.Contains(err.Error(), "Username and Password not accepted") {
			fmt.Println("Invalid username and password")
			return []domain.EmailResult{
				{
					Email:   user.Email,
					Status:  "failed",
					Message: err.Error(),
				},
			}
		} else {
			results = append(results, domain.EmailResult{
				Email:   user.Email,
				Status:  "failed",
				Message: err.Error(),
			})
		}
	} else {
		results = append(results, domain.EmailResult{
			Email:   user.Email,
			Status:  "sent",
			Message: "Email sent successfully",
		})

	}

	var mu sync.Mutex
	// Store results for each email

	wg := sync.WaitGroup{}
	emailJobs := make(chan domain.User, len(users)-1)

	// Worker pool
	for i := 0; i < maxWorkers; i++ {
		go func() {
			for user := range emailJobs {
				err := sender.Send(user.Email, user.Name)
				mu.Lock()
				if err != nil {
					results = append(results, domain.EmailResult{
						Email:   user.Email,
						Status:  "failed",
						Message: err.Error(),
					})
				} else {
					results = append(results, domain.EmailResult{
						Email:   user.Email,
						Status:  "sent",
						Message: "Email sent successfully",
					})
				}
				mu.Unlock()
				wg.Done()
			}
		}()
	}

	// Add jobs to worker pool starting from index 1
	for _, user := range users[1:] {
		wg.Add(1)
		emailJobs <- user
	}
	close(emailJobs)
	wg.Wait()

	return results // Return results for all emails
}
