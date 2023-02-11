package service

import (
	"github.com/eduaravila/momo/apps/auth/internal/adapter"
	"github.com/eduaravila/momo/apps/auth/internal/factory"
	"github.com/eduaravila/momo/apps/auth/internal/storage"
	"github.com/eduaravila/momo/apps/auth/internal/types"
	"github.com/eduaravila/momo/packages/db/queries"
)

type LoginOIDCService struct {
	storage Storage
	api     AuthApi
}

type AuthApi interface {
	GetToken(code string) (*types.OAuthToken, error)
	GetOidcUserInfo(*types.OAuthToken) (*types.OIDCClaims, error)
}

type Storage interface {
	CreateSession(queries.Session) (queries.Session, error)
	CreateUserAccount(types.OIDCClaims, types.OAuthToken) (*storage.UserAccount, error)
}

func NewTwitchHandler(storage Storage, api *adapter.TwitchAPI) *LoginOIDCService {
	return &LoginOIDCService{storage: storage, api: api}
}

func (t *LoginOIDCService) LogIn(code string, metadata types.Metadata) (string, error) {

	tokenResponse, err := t.api.GetToken(code)

	if err != nil {
		return "", err
	}
	userInfoResponse, err := t.api.GetOidcUserInfo(tokenResponse)
	if err != nil {
		return "", err
	}

	userAccount, err := t.storage.CreateUserAccount(*userInfoResponse, *tokenResponse)

	if err != nil {
		return "", err
	}

	token := factory.NewSessionToken(userAccount.User.ID.String())
	tokenString, err := token.Sign()
	if err != nil {
		return "", err
	}

	t.storage.CreateSession(queries.Session{
		ExpiredAt:    token.Claims().ExpiresAt.Time,
		UserAgent:    metadata.UserAgent,
		UserID:       userAccount.User.ID,
		SessionToken: tokenString,
		IpAddress:    metadata.IPAddress,
	})

	return tokenString, nil
}
