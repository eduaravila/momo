package service

import (
	"github.com/eduaravila/momo/apps/auth/internal/adapter"
	"github.com/eduaravila/momo/apps/auth/internal/app"
	"github.com/eduaravila/momo/apps/auth/internal/app/command"
	"github.com/eduaravila/momo/apps/auth/internal/storage"
	"github.com/eduaravila/momo/packages/postgres/queries"
)

func NewApplication() app.Application {

	tokenService := adapter.NewJwtTokenCreator()
	oAuthService := adapter.NewTwitchAPI()

	return newApplication(oAuthService, tokenService)
}

func newApplication(
	oAuthService command.OAuthService,
	tokenService command.TokenService,
) app.Application {
	postgreDB, err := storage.InitPostgresDB()
	if err != nil {
		panic(err)
	}
	queries := queries.New(postgreDB)
	sessionStorage := storage.NewSessionPostgresStorage(queries)

	return app.Application{
		Queries: app.Queries{},
		Commands: app.Commands{
			AuthenticateWithOIDC: command.NewAuthenticateWithOIDCHandler(oAuthService, tokenService, sessionStorage),
		},
	}
}
