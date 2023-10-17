package providers

import (
	"context"
	"email-service/pkg/mtools/mconfig"
	"log"

	"github.com/mailgun/mailgun-go/v4"
)

// MailgunProvider represents a provider for sending emails using Mailgun service.
type MailgunProvider struct{}

// SendEmail sends an email using MailgunProvider.
func (p *MailgunProvider) SendEmail(to, subject, body string) error {
	env := mconfig.GlobalEnv
	mg := mailgun.NewMailgun(env.Domain, env.ApiKeyMailGun)

	message := mg.NewMessage(
		env.Sender,
		subject,
		body,
		to,
	)

	ctx := context.Background()
	_, _, err := mg.Send(ctx, message)
	if err != nil {
		log.Printf("Error sending email via Mailgun: %v", err)
		return err
	}
	return nil
}

// NewMailgunProvider creates a new instance of MailgunProvider.
func NewMailgunProvider() *MailgunProvider {
	return &MailgunProvider{}
}
