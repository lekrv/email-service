package domain

// EmailRepository defines the interface for sending emails using different providers.
type EmailRepository interface {
	SendEmail(provider string, email Email) error
}

// Email represents an email message with a subject, body, and recipient email address.
type Email struct {
	From    string `json:"from"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
	ToEmail string `json:"to_email"`
}
