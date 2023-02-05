package factory

import (
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type SessionToken struct {
	Token *jwt.Token
}

func DefaultClaimsForSession(id string) jwt.RegisteredClaims {
	exp := time.Now().Add(1 * time.Hour)

	return jwt.RegisteredClaims{
		Subject:   id,
		ExpiresAt: jwt.NewNumericDate(exp),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
}

func (t *SessionToken) Claims() jwt.RegisteredClaims {
	return t.Token.Claims.(jwt.RegisteredClaims)
}

func NewSessionToken(subject string) *SessionToken {
	return &SessionToken{
		jwt.NewWithClaims(jwt.SigningMethodES256, DefaultClaimsForSession(subject)),
	}
}

func (t *SessionToken) SetClaims(claims jwt.RegisteredClaims) *SessionToken {
	t.Token.Claims = claims
	return t
}

func (t *SessionToken) Sign() (string, error) {
	decodedKey, err := jwt.ParseECPrivateKeyFromPEM([]byte(os.Getenv("JWT_PRIVATE_KEY")))
	if err != nil {
		return "", err
	}
	return t.Token.SignedString(decodedKey)
}
