package command

import (
	"context"

	"github.com/eduaravila/momo/apps/auth/internal/oauth/domain/session"
	"github.com/eduaravila/momo/packages/decorators"
)

/*
GenerateSession receives OIDC code and scope to generate a session token
** this will be displayed in the logs and metrics **
*/
type GenerateSession struct {
	code  string
	scope []string
}

type generateSessionTokenHandler struct {
	oAuthRepository OAuthService
	tokenService    TokenService
	repo            session.Storage
}

func NewGenerateSessionTokenHandler(oAuthRepository OAuthService, repo session.Storage) decorators.CommandHandler[GenerateSession] {
	return &generateSessionTokenHandler{
		oAuthRepository: oAuthRepository,
		repo:            repo,
	}
}

func (g *generateSessionTokenHandler) Handle(ctx context.Context, cmd GenerateSession) error {
	//TODOA
}
