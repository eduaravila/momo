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
}

type generateSessionTokenHandler struct {
	oAuthService OAuthService
	tokenService TokenService
	repo         session.Storage
}

func NewGenerateSessionTokenHandler(oAuthRepository OAuthService, repo session.Storage) decorators.CommandHandler[GenerateSession] {
	return &generateSessionTokenHandler{
		oAuthService: oAuthRepository,
		repo:         repo,
	}
}

func (g *generateSessionTokenHandler) Handle(ctx context.Context, cmd GenerateSession) error {
	accessInfo, err := g.oAuthService.GetAuthorizationInformation(ctx, cmd.code)
	if err != nil {
		return errors.Join(err, errors.New("failed to get token"))
	}
	claims, err := g.oAuthService.GetOIDCUserInfo(ctx, accessInfo.AccessToken)


	if err != nil {
		return errors.Join(err, errors.New("failed to get user info"))
	}

	userAccount, err := g..CreateUserAccount(*userInfoResponse, *tokenResponse)
	if err != nil {

		return nil, err
	}

	token := factory.NewSessionToken(userAccount.User.ID.String())
	tokenString, err := token.Sign()

	if err != nil {
		return nil, err
	}

	session, err := t.storage.CreateSession(queries.Session{
		ExpiredAt:    token.Claims().ExpiresAt.Time,
		UserAgent:    metadata.UserAgent,
		UserID:       userAccount.User.ID,
		SessionToken: tokenString,
		IpAddress:    metadata.IPAddress,
	})

}
