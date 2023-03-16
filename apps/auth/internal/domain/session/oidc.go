package session

import (
	"time"
)

type OIDCAccount struct {
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

type OIDCAccessToken struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int
	TokenType    string
	Scope        []string
}

func UnmarshalOIDCAccessTokenFromDatabase(
	accessToken string,
	refreshToken string,
	expiresIn int,
	tokenType string,
	scope []string,
) *OIDCAccessToken {
	return &OIDCAccessToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
		TokenType:    tokenType,
		Scope:        scope,
	}
}

func NewOIDCAccount(
	aud string,
	exp int64,
	iat int64,
	iss string,
	sub string,
	email string,
	emailVerified bool,
	picture string,
	preferedUsername string,
	updatedAt time.Time,
) (*OIDCAccount, error) {
	// TODO: Validate input
	return &OIDCAccount{
		Aud:              aud,
		Exp:              exp,
		Iat:              iat,
		Iss:              iss,
		Sub:              sub,
		Email:            email,
		EmailVerified:    emailVerified,
		Picture:          picture,
		PreferedUsername: preferedUsername,
		UpdatedAt:        updatedAt,
	}, nil
}
