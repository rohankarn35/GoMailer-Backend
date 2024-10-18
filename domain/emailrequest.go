package domain

type EmailRequest struct {
	Users        []User `json:"users"`
	HTMLTemplate string `json:"html_template"`
	SenderEmail  string `json:"sender_email"`
	SenderName   string `json:"sender_name"`
	AppPassword  string `json:"app_password"`
	Subject      string `json:"subject"`
}
