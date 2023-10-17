package providers

import (
	"errors"
	"log"
	"net/http"

	"email-service/pkg/mtools/mconfig"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendGridProvider is a struct that contains the SendGrid API key.
type SendGridProvider struct{}

// SendEmail sends an email using the SendGrid API.
func (p *SendGridProvider) SendEmail(toEmail, emailSubject string, emailBody string) error {
	env := mconfig.GlobalEnv
	client := sendgrid.NewSendClient(env.ApiKeySendgrid)

	from := mail.NewEmail(env.UserNameSendgrid, env.FromEmailSendgrid)
	subject := emailSubject
	to := mail.NewEmail(toEmail, toEmail)
	content := mail.NewContent("text/html", emailBody)
	message := mail.NewV3MailInit(from, subject, to, content)

	response, err := client.Send(message)
	if err != nil {
		log.Printf("Error sending email via SendGrid: %v", err)
		return err
	}

	if response.StatusCode == http.StatusBadRequest {
		log.Printf("Error sending email via SendGrid: %v", err)
		return errors.New(response.Body)
	}

	return nil
}

// NewSendGridProvider creates a new SendGridProvider instance with the given API key.
func NewSendGridProvider() *SendGridProvider {
	return &SendGridProvider{}
}
