package command

import "context"

type OAuthService interface {
	GetToken(ctx context.Context, code string) (string, error)
	GetOIDCUserInfo(ctx context.Context, code string) (string, error)
}

type TokenService interface {
	CreateSessionToken(ctx context.Context, subject string) (string, error)
}
