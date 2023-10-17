//go:build wireinject
// +build wireinject

package di

import (
	"email-service/internal/infrastructure/server/handlers/emails"
	usecase "email-service/internal/usecases"
	"email-service/pkg/mtools/mjsonapi"

	"github.com/google/wire"
)

// stdSet is a Wire provider set that includes providers for creating instances of:
// - ErrorsJSONAPIProvider
// - ResponseJSONAPI
// - Logger
// - EmailRepository
// - EmailUseCase
// - EmailAction
var stdSet = wire.NewSet(
	mjsonapi.NewErrorsJSONAPI,
	mjsonapi.NewResponseJSONAPI,

	loggerProvider,
	sendGridProvider,
	mailgunProvider,
	emailRepositoryProvider,

	usecase.NewEmailUseCase,
	emails.NewHandler,

	wire.Bind(new(emails.EmailAction), new(*usecase.EmailUseCase)),
	wire.Bind(new(mjsonapi.ErrorsJSONAPIProvider), new(*mjsonapi.ErrorsJSONAPI)),
)
