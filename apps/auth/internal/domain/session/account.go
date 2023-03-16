package session

import (
	"errors"
	"time"
)

type Account struct {
	ID           string
	UserID       string
	AccessToken  string
	RefreshToken string

	Picture          string
	Email            string
	PreferedUsername string

	Iss   string
	Sub   string
	Scope []string

	ExpiredAt time.Time
}

func NewAccount(
	id string,
	userID string,
	accessToken string,
	refreshToken string,
	picture string,
	email string,
	preferedUsername string,
	iss string,
	sub string,
	scope []string,
) (*Account, error) {
	if id == "" {
		return nil, errors.New("no id provided")
	}

	if userID == "" {
		return nil, errors.New("no user id provided")
	}

	if accessToken == "" {
		return nil, errors.New("no access token provided")
	}

	if refreshToken == "" {
		return nil, errors.New("no refresh token provided")
	}

	if email == "" {
		return nil, errors.New("no email provided")
	}

	if iss == "" {
		return nil, errors.New("no issuer provided")
	}

	if sub == "" {
		return nil, errors.New("no subject provided")
	}

	if len(scope) == 0 {
		return nil, errors.New("no scope provided")
	}

	return &Account{
		ID:               id,
		UserID:           userID,
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		Picture:          picture,
		Email:            email,
		PreferedUsername: preferedUsername,
		Iss:              iss,
		Sub:              sub,
		Scope:            scope,
	}, nil
}

func NewAccountFromOIDC(
	accountUUID,
	userUUID string,
	account *OIDCAccount,
	token *OIDCAccessToken,
) (*Account, error) {
	if account == nil {
		return nil, errors.New("no account provided")
	}

	return NewAccount(
		accountUUID,
		userUUID,
		token.AccessToken,
		token.RefreshToken,
		account.Picture,
		account.Email,
		account.PreferedUsername,
		account.Iss,
		account.Sub,
		token.Scope,
	)
}
