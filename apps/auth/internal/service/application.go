package service

import (
	"os"

	"github.com/eduaravila/momo/apps/auth/internal/adapter"
	"github.com/eduaravila/momo/apps/auth/internal/app"
	"github.com/eduaravila/momo/apps/auth/internal/app/command"
	"github.com/eduaravila/momo/apps/auth/internal/app/query"
	"github.com/eduaravila/momo/apps/auth/internal/storage"
	"github.com/eduaravila/momo/packages/metrics"
	"github.com/eduaravila/momo/packages/postgres/queries"
	"golang.org/x/exp/slog"
)

func NewApplication() app.Application {
	tokenService := adapter.NewJwtTokenCreator()
	oAuthService := adapter.NewTwitchAPI()

	return newApplication(oAuthService, tokenService, tokenService)
}

func NewTestApplication() app.Application {
	tokenService := &TokenServiceMock{}

	return newApplication(&OAuthServiceMock{}, tokenService, tokenService)
}

func newApplication(
	oAuthService command.OAuthService,
	tokenService command.TokenService,
	tokenVerifierService query.TokenVeriferService,
) app.Application {
	postgreDB, err := storage.InitPostgresDB()
	if err != nil {
		panic(err)
	}

	queries := queries.New(postgreDB)

	sessionStorage := storage.NewSessionPostgresStorage(queries)
	logger := slog.New(slog.NewTextHandler(os.Stdin))
	metricsClient := metrics.NoOpt{}

	return app.Application{
		Queries: app.Queries{
			SessionWithID: query.NewSessionWithIDHandler(sessionStorage),
			VerifySessionToken: query.NewSessionTokenVerifierHandler(
				tokenVerifierService,
				logger,
				metricsClient,
			),
		},
		Commands: app.Commands{
			AuthenticateWithOIDC: command.NewAuthenticateWithOIDCHandler(
				oAuthService,
				tokenService,
				sessionStorage,
				logger,
				metricsClient,
			),
		},
	}
}
