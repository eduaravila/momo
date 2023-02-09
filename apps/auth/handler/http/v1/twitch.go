package v1

import (
	"github.com/eduaravila/momo/apps/auth/factory"
	"github.com/eduaravila/momo/apps/auth/storage"
	"github.com/eduaravila/momo/apps/auth/svc"
	"github.com/eduaravila/momo/apps/auth/types"
	"github.com/eduaravila/momo/packages/db/queries"
)

type TwitchHandler struct {
	storage Storage
	api     *svc.TwitchAPI
}

type Storage interface {
	CreateSession(queries.Session) (queries.Session, error)
	CreateUserAccount(*types.UserinfoRespose, *types.TokenResponse) (*storage.UserAccountSession, error)
}

type Metadata struct {
	UserAgent string
	IPAddress string
}

func NewTwitchHandler(storage Storage, api *svc.TwitchAPI) *TwitchHandler {
	return &TwitchHandler{storage: storage, api: api}
}

func (t *TwitchHandler) LogIn(code string, metadata Metadata) (string, error) {

	oidcToken, err := t.api.GetToken(code)

	if err != nil {
		return "", err
	}
	userInfo, err := t.api.GetOidcUserInfo(oidcToken)
	if err != nil {
		return "", err
	}

	userAndAccount, err := t.storage.CreateUserAccount(userInfo, oidcToken)

	if err != nil {
		return "", err
	}

	token := factory.NewSessionToken(userAndAccount.User.ID.String())
	tokenString, err := token.Sign()
	if err != nil {
		return "", err
	}

	t.storage.CreateSession(queries.Session{
		ExpiredAt:    token.Claims().ExpiresAt.Time,
		UserAgent:    metadata.UserAgent,
		UserID:       userAndAccount.User.ID,
		SessionToken: tokenString,
		IpAddress:    metadata.IPAddress,
	})

	return tokenString, nil
}
