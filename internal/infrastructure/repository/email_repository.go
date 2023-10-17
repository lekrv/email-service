package repository

import (
	"errors"

	"email-service/cmd/di/providers"
	"email-service/internal/domain"
)

// EmailRepository represents a repository for sending emails using different providers.
type EmailRepository struct {
	SendGridService *providers.SendGridProvider
	MailgunService  *providers.MailgunProvider
}

// SendEmail sends an email using the specified provider.
func (er *EmailRepository) SendEmail(provider string, email domain.Email) error {
	switch provider {
	case "sendgrid":
		return er.SendGridService.SendEmail(email.ToEmail, email.Subject, email.Body)
	case "mailgun":
		return er.MailgunService.SendEmail(email.ToEmail, email.Subject, email.Body)
	default:
		return errors.New("unknown provider")
	}
}

// NewEmailRepository creates a new instance of EmailRepository.
func NewEmailRepository(
	sendGridService *providers.SendGridProvider,
	mailgunService *providers.MailgunProvider,
) *EmailRepository {
	return &EmailRepository{
		SendGridService: sendGridService,
		MailgunService:  mailgunService,
	}
}
