package adapter

import (
	"errors"
	"os"
	"time"

	"github.com/eduaravila/momo/apps/auth/internal/domain/session"

	jwt "github.com/golang-jwt/jwt/v4"
)

type SessionToken struct {
	Claims jwt.RegisteredClaims
}

type JWTToken struct {
	token *jwt.Token
}

func DefaultClaimsForSessionInclude(id string) jwt.RegisteredClaims {
	exp := time.Now().Add(1 * time.Hour)

	return jwt.RegisteredClaims{
		Subject: id,

		ExpiresAt: jwt.NewNumericDate(exp),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
}

func NewJwtTokenCreator(claims jwt.RegisteredClaims) *SessionToken {
	return &SessionToken{
		claims,
	}
}

func (s *SessionToken) CreateSessionToken(subject string) (session.Token, error) {
	claims := DefaultClaimsForSessionInclude(subject)
	token := NewJWTToken(DefaultClaimsForSessionInclude((subject)))
	signedToken, err := token.Sign()

	if err != nil {
		return session.Token{}, errors.Join(err, errors.New("failed creating session token"))
	}

	return session.NewSessionToken(signedToken, true,
		session.NewClaims(
			claims.Issuer,
			claims.Subject,
			claims.Audience,
			claims.ExpiresAt.Time,
			claims.NotBefore.Time,
			claims.IssuedAt.Time,
			claims.ID)), nil
}

func NewJWTToken(claims jwt.RegisteredClaims) *JWTToken {
	return &JWTToken{
		token: jwt.NewWithClaims(jwt.SigningMethodES256, claims),
	}
}

func (t *JWTToken) Sign() (string, error) {
	decodedKey, err := jwt.ParseECPrivateKeyFromPEM([]byte(os.Getenv("JWT_PRIVATE_KEY")))
	if err != nil {
		return "", err
	}

	return t.token.SignedString(decodedKey)
}

func (t *JWTToken) Verify(tokenString string) error {
	decodedKey, err := jwt.ParseECPublicKeyFromPEM([]byte(os.Getenv("JWT_PUBLIC_KEY")))
	if err != nil {
		return err
	}

	_, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return decodedKey, nil
	})

	return err
}
