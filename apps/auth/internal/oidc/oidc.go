package oidc

import (
	"github.com/eduaravila/momo/apps/auth/internal/adapter"
	"github.com/eduaravila/momo/apps/auth/internal/factory"
	"github.com/eduaravila/momo/apps/auth/internal/types"
	"github.com/eduaravila/momo/packages/db/queries"
)
type Storage interface {
	CreateSession(ctx, context.Context, .Session) (queries.Session, error)
	CreateUserAccount(types.OIDCClaims, types.OAuthToken) (*storage.UserAccount, error)
}





type OIDC struct {
	code string
}

type LoginOIDCService struct {
	storage Storage
	api     AuthAPI
}

func NewAuthService(storage Storage, api *adapter.TwitchAPI) *LoginOIDCService {
	return &LoginOIDCService{storage: storage, api: api}
}

func (t *LoginOIDCService) LogIn(code string, metadata types.Metadata) (*queries.Session, error) {
	tokenResponse, err := t.api.GetToken(code)
	if err != nil {
		return nil, err
	}

	userInfoResponse, err := t.api.GetOidcUserInfo(tokenResponse)
	if err != nil {
		return nil, err
	}

	userAccount, err := t.storage.CreateUserAccount(*userInfoResponse, *tokenResponse)
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

	return &session, err
}
