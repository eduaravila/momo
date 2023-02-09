package service

import (
	"database/sql"
	"strings"
	"time"

	"github.com/eduaravila/momo/apps/auth/factory"
	"github.com/eduaravila/momo/apps/auth/svc"
	"github.com/eduaravila/momo/apps/auth/types"
	"github.com/eduaravila/momo/packages/db/queries"
	"github.com/google/uuid"
)

type TwitchHandler struct {
	storage Storage
	api     *svc.TwitchAPI
}

type Storage interface {
	CreateSession(queries.Session) (queries.Session, error)
	GetAccountAndUserBySub(string) (queries.GetAccountAndUserBySubRow, error)
	CreateUser() (queries.User, error)
	CreateAccount(queries.Account) (queries.Account, error)
}

func NewTwitchHandler(storage Storage, api *svc.TwitchAPI) *TwitchHandler {
	return &TwitchHandler{storage: storage, api: api}
}

func (t *TwitchHandler) LogIn(code string, metadata types.Metadata) (string, error) {

	tokenResponse, err := t.api.GetToken(code)

	if err != nil {
		return "", err
	}
	userInfoResponse, err := t.api.GetOidcUserInfo(tokenResponse)
	if err != nil {
		return "", err
	}

	result, err := t.storage.GetAccountAndUserBySub(userInfoResponse.Sub)
	uid := result.UserID
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	var user queries.User
	if err == sql.ErrNoRows {
		user, err = t.storage.CreateUser()
		if err != nil {
			return "", err
		}
	}

	_, err = t.storage.CreateAccount(queries.Account{
		AccessToken:      tokenResponse.AccessToken,
		RefreshToken:     tokenResponse.RefreshToken,
		ExpiredAt:        time.Unix(int64(tokenResponse.ExpiresIn), 0),
		Email:            userInfoResponse.Email,
		Picture:          userInfoResponse.Picture,
		Iss:              userInfoResponse.Iss,
		PreferedUsername: userInfoResponse.PreferedUsername,
		ID:               uuid.New(),
		Scope:            strings.Join(tokenResponse.Scope, " "),
		Sub:              userInfoResponse.Sub,
		UserID:           uid,
	})
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	token := factory.NewSessionToken(user.ID.String())
	tokenString, err := token.Sign()
	if err != nil {
		return "", err
	}

	t.storage.CreateSession(queries.Session{
		ExpiredAt:    token.Claims().ExpiresAt.Time,
		UserAgent:    metadata.UserAgent,
		UserID:       user.ID,
		SessionToken: tokenString,
		IpAddress:    metadata.IPAddress,
	})

	return tokenString, nil
}
