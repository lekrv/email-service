//go:build wireinject
// +build wireinject

package di

import (
	"email-service/internal/infrastructure/server/handlers/emails"
	"github.com/google/wire"
)

func Initialize() (*emails.Handler, error) {
	wire.Build(stdSet)

	return &emails.Handler{}, nil
}
