package command

import (
	"context"
	"errors"

	"github.com/eduaravila/momo/apps/auth/internal/oauth/domain/session"
	"github.com/eduaravila/momo/packages/decorators"
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

type AuthenticateWithOIDCHandler decorators.CommandHandler[authenticateWithOIDCHandler]

func NewAuthenticateWithOIDCHandler(
	oAuthRepository OAuthService,
	store session.Storage,
) decorators.CommandHandler[GenerateSession] {
	return &authenticateWithOIDCHandler{
		oAuthService: oAuthRepository,
		store:        store,
	}
}

func (g *authenticateWithOIDCHandler) Handle(
	ctx context.Context,
	cmd GenerateSession,
) error {
	account, err := g.oAuthService.GetAccount(ctx, cmd.Code, cmd.AccountUUID, cmd.UserUUID)
	if err != nil {
		return errors.Join(err, errors.New("failed to get token"))
	}

	err = g.store.AddAccountWithUser(ctx, account, cmd.UserUUID)
	if err != nil {
		return errors.Join(err, errors.New("failed to add account"))
	}

	token, err := g.tokenService.CreateSessionToken(ctx, account.UserID)

	if err != nil {
		return err
	}

	session := session.NewSession(
		cmd.SessionUUID,
		cmd.UserUUID,
		token.Claims.ExpiresAt,
		token,
		cmd.Metadata.UserAgent,
		cmd.Metadata.IPAddress,
	)

	return g.store.AddSession(ctx, session)
}
