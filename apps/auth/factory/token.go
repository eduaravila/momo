package factory

import (
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type Token struct {
	Claims jwt.RegisteredClaims
}

func DefaultClaimsForSession(id string) jwt.RegisteredClaims {
	exp := time.Now().Add(1 * time.Hour)

	return jwt.RegisteredClaims{
		Subject:   id,
		ExpiresAt: jwt.NewNumericDate(exp),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
}

func NewToken() *Token {
	return &Token{}
}

func (t *Token) SetClaims(claims jwt.RegisteredClaims) *Token {
	t.Claims = claims
	return t
}

func (t *Token) Generate() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, t.Claims)
	decodedKey, err := jwt.ParseECPrivateKeyFromPEM([]byte(os.Getenv("JWT_PRIVATE_KEY")))
	if err != nil {
		return "", err
	}

	return token.SignedString(decodedKey)
}
