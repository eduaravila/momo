package oidc

import "github.com/eduaravila/momo/apps/auth/internal/types"

type Storage interface {
	GetToken(code string) (*types.OAuthToken, error)
	GetOidcUserInfo(*types.OAuthToken) (*types.OIDCClaims, error)
}
