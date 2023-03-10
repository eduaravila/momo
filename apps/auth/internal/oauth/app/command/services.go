package command

import (
	"context"

	"github.com/eduaravila/momo/apps/auth/internal/oauth/domain/session"
)

type OAuthService interface {
	GetAuthorizationInformation(ctx context.Context, code string) (session.OIDCOAuthToken, error)
	GetOIDCUserInfo(ctx context.Context, accessToken string) (session.OIDCClaims, error)
}

type TokenService interface {
	CreateSessionToken(ctx context.Context, subject string) (session.Token, error)
}
