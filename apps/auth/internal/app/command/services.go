package command

import (
	"context"

	"github.com/eduaravila/momo/apps/auth/internal/domain/session"
)

type OAuthService interface {
	GetAccount(ctx context.Context, accessToken, accountUUID, userUUID string) (*session.Account, error)
}

type TokenService interface {
	CreateSessionToken(ctx context.Context, subject string) (*session.Token, error)
}
