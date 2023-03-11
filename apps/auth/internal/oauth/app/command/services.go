package command

import (
	"context"

	"github.com/eduaravila/momo/apps/auth/internal/oauth/domain/session"
)

type OAuthService interface {
	GetAccount(ctx context.Context, accessToken string) (*session.Account, error)
}

type TokenService interface {
	CreateSessionToken(ctx context.Context, subject string) (session.Token, error)
}
