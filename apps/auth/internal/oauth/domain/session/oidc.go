package session

import (
	"time"

	"github.com/eduaravila/momo/apps/auth/internal/adapter"
	"github.com/eduaravila/momo/apps/auth/internal/factory"
	"github.com/eduaravila/momo/apps/auth/internal/types"
	"github.com/eduaravila/momo/packages/db/queries"
)

type OIDCClaims struct {
	Aud              string
	Exp              int64
	Iat              int64
	Iss              string
	Sub              string
	Email            string
	EmailVerified    bool
	Picture          string
	PreferedUsername string
	UpdatedAt        time.Time
}

type OIDCOAuthToken struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int
	TokenType    string
	Scope        []string
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
