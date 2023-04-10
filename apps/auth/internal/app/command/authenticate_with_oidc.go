package command

import (
	"context"
	"errors"

	"github.com/eduaravila/momo/apps/auth/internal/domain/session"
	"github.com/eduaravila/momo/packages/decorators"
	"golang.org/x/exp/slog"
)

/*
GenerateSession receives OIDC code and scope to generate a session token
** this will be displayed in the logs and metrics **
*/
type GenerateSession struct {
	Code        string
	Scope       []string
	SessionUUID string
	AccountUUID string
	UserUUID    string
	Metadata    session.ClientMetadata
}

type authenticateWithOIDCHandler struct {
	oAuthService OAuthService
	tokenService TokenService
	store        session.Storage
}

type AuthenticateWithOIDCHandler decorators.CommandHandler[GenerateSession]

func NewAuthenticateWithOIDCHandler(
	oAuthRepository OAuthService,
	tokenService TokenService,
	store session.Storage,
	logger *slog.Logger,
	metrics decorators.MetricsClient,
) AuthenticateWithOIDCHandler {
	return decorators.ApplyCommandDecorators[GenerateSession](
		&authenticateWithOIDCHandler{
			oAuthService: oAuthRepository,
			tokenService: tokenService,
			store:        store,
		},
		logger,
		metrics,
	)
}

func (g *authenticateWithOIDCHandler) Handle(
	ctx context.Context,
	cmd GenerateSession,
) error {
	account, err := g.oAuthService.GetAccount(ctx, cmd.Code, cmd.AccountUUID, cmd.UserUUID)

	if err != nil {
		return errors.Join(err, errors.New("failed to get token"))
	}

	err = g.store.AddAccountWithUser(ctx, account)
	if err != nil {
		return errors.Join(err, errors.New("failed to add account"))
	}

	token, err := g.tokenService.CreateSessionToken(ctx, account.UserID)

	if err != nil {
		return err
	}

	session, err := session.NewSession(
		cmd.SessionUUID,
		account.UserID,
		token.Claims.ExpiresAt,
		token,
		cmd.Metadata.UserAgent,
		cmd.Metadata.IPAddress,
	)

	if err != nil {
		return err
	}

	return g.store.AddSession(ctx, session)
}
