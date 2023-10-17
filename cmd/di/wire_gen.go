// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"email-service/internal/infrastructure/server/handlers/emails"
	"email-service/internal/usecases"
	"email-service/pkg/mtools/mjsonapi"
)

// Injectors from wire.go:

func Initialize() (*emails.Handler, error) {
	providersSendGridProvider := sendGridProvider()
	providersMailgunProvider := mailgunProvider()
	emailRepository := emailRepositoryProvider(providersSendGridProvider, providersMailgunProvider)
	emailUseCase := usecase.NewEmailUseCase(emailRepository)
	errorsJSONAPI := mjsonapi.NewErrorsJSONAPI()
	responseJSONAPIProvider := mjsonapi.NewResponseJSONAPI()
	logger := loggerProvider()
	handler := emails.NewHandler(emailUseCase, errorsJSONAPI, responseJSONAPIProvider, logger)
	return handler, nil
}
