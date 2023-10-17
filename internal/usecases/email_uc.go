package usecase

import (
	"log"

	"email-service/internal/domain"
)

const (
	sendGridProvider string = "sendgrid"
	mailgunProvider  string = "mailgun"
)

// EmailUseCase represents the use case for sending emails.
type EmailUseCase struct {
	emailRepository domain.EmailRepository
}

// SendEmail sends an email using the email repository.
// If the email fails to send via SendGrid, it will try to send it via Mailgun.
// If both providers fail, it will return an error.
func (uc *EmailUseCase) SendEmail(email domain.Email) error {
	var err error

	if err := uc.emailRepository.SendEmail(sendGridProvider, email); err == nil {
		log.Printf("Email delivered via SendGrid")
		return nil
	}

	log.Printf("Error sending email via SendGrid: %v", err)

	if err := uc.emailRepository.SendEmail(mailgunProvider, email); err != nil {
		log.Printf("Error sending email via Mailgun: %v", err)
		log.Printf("Email delivery failed via both providers")
		return err
	}

	log.Printf("Email delivered via Mailgun")
	return nil
}

// NewEmailUseCase creates a new EmailUseCase instance.
func NewEmailUseCase(emailRepository domain.EmailRepository) *EmailUseCase {
	return &EmailUseCase{
		emailRepository: emailRepository,
	}
}
