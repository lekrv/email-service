package di

import (
	"email-service/cmd/di/providers"
	"email-service/internal/domain"
	"email-service/internal/infrastructure/repository"
	"email-service/pkg/mtools/mlogger"
)

// loggerProvider returns a new instance of mlogger.Logger using LogrusProvider.
func loggerProvider() mlogger.Logger {
	return mlogger.NewLogrusProvider().Logger()
}

// SendGridProvider returns a new instance of SendGridProvider with the given API key.
func sendGridProvider() *providers.SendGridProvider {
	return providers.NewSendGridProvider()
}

// mailgunProvider returns a new instance of MailgunProvider with the given API key, domain and sender.
func mailgunProvider() *providers.MailgunProvider {
	return providers.NewMailgunProvider()
}

// emailRepositoryProvider returns a new instance of domain.EmailRepository using SendGridProvider and MailgunProvider.
func emailRepositoryProvider(sendGridService *providers.SendGridProvider, mailgunService *providers.MailgunProvider) domain.EmailRepository {
	return repository.NewEmailRepository(sendGridService, mailgunService)
}
