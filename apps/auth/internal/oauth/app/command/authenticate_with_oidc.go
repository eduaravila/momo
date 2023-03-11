package command

import (
	"context"
	"errors"

	"github.com/eduaravila/momo/apps/auth/internal/oauth/domain/session"
	"github.com/eduaravila/momo/packages/db/queries"
	"github.com/eduaravila/momo/packages/decorators"
)

/*
GenerateSession receives OIDC code and scope to generate a session token
** this will be displayed in the logs and metrics **
*/
type GenerateSession struct {
	code        string
	scope       []string
	SessionUUID string
	AccountUUID string
}

type authenticateWithOIDCHandler struct {
	oAuthService OAuthService
	tokenService TokenService
	repo         session.Storage
}

type AuthenticateWithOIDCHandler decorators.CommandHandler[authenticateWithOIDCHandler]

func NewAuthenticateWithOIDCHandler(oAuthRepository OAuthService, repo session.Storage) decorators.CommandHandler[GenerateSession] {
	return &authenticateWithOIDCHandler{
		oAuthService: oAuthRepository,
		repo:         repo,
	}
}

func (g *authenticateWithOIDCHandler) Handle(ctx context.Context, cmd GenerateSession) error {
	accessInfo, err := g.oAuthService.GetAuthorizationInformation(ctx, cmd.code)
	if err != nil {
		return errors.Join(err, errors.New("failed to get token"))
	}
	claims, err := g.oAuthService.GetOIDCAccount(ctx, accessInfo.AccessToken)

	if err != nil {
		return errors.Join(err, errors.New("failed to get user info"))
	}

	user, err := g.repo.FindUserFromSub(ctx, claims.Sub)

	account, err := session.NewAccount(cmd.AccountUUID, user.ID, accessInfo.AccessToken, accessInfo.RefreshToken, claims.Picture, claims.Email, claims.PreferedUsername, claims.Iss, claims.Sub, accessInfo.Scope)

	err = g.repo.AddAccountWithUser(ctx, account)
	if err != nil {
		return errors.Join(err, errors.New("failed to add account"))
	}

	token := factory.NewSessionToken(userAccount.User.ID.String())
	tokenString, err := token.Sign()

	if err != nil {
		return err
	}

	session, err := t.storage.CreateSession(queries.Session{
		ExpiredAt:    token.Claims().ExpiresAt.Time,
		UserAgent:    metadata.UserAgent,
		UserID:       userAccount.User.ID,
		SessionToken: tokenString,
		IpAddress:    metadata.IPAddress,
	})

}
