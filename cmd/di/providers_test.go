package di

import (
	"email-service/cmd/di/providers"
	"email-service/internal/domain"
	_ "email-service/internal/infrastructure/repository"
	_ "email-service/pkg/mtools/mlogger"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoggerProvider(t *testing.T) {
	// Call the loggerProvider function.
	logger := loggerProvider()

	// Assert that the returned value is not nil.
	assert.NotNil(t, logger)
}

func TestSendGridProvider(t *testing.T) {
	// Call the sendGridProvider function.
	provider := sendGridProvider()

	// Assert that the returned value is of type *providers.SendGridProvider.
	assert.IsType(t, &providers.SendGridProvider{}, provider)
}

func TestMailgunProvider(t *testing.T) {
	// Call the mailgunProvider function.
	provider := mailgunProvider()

	// Assert that the returned value is of type *providers.MailgunProvider.
	assert.IsType(t, &providers.MailgunProvider{}, provider)
}

func TestEmailRepositoryProvider(t *testing.T) {
	// Create mock instances of SendGridProvider and MailgunProvider.
	sendGridProvider := &providers.SendGridProvider{}
	mailgunProvider := &providers.MailgunProvider{}

	// Call the emailRepositoryProvider function.
	repository := emailRepositoryProvider(sendGridProvider, mailgunProvider)

	// Assert that the returned value implements the domain.EmailRepository interface.
	assert.Implements(t, (*domain.EmailRepository)(nil), repository)
}
