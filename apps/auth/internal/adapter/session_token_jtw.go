package adapter

import (
	"context"
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

func NewJwtTokenCreator() *SessionToken {
	return &SessionToken{}
}

func (s *SessionToken) CreateSessionToken(ctx context.Context, subject string) (*session.Token, error) {
	claims := DefaultClaimsForSessionInclude(subject)
	token := NewJWTToken(DefaultClaimsForSessionInclude((subject)))
	signedToken, err := token.Sign()

	if err != nil {
		return nil, errors.Join(err, errors.New("failed creating session token"))
	}

	return session.NewSessionToken(signedToken, true,
		session.NewClaims(
			claims.Issuer,
			claims.Subject,
			claims.Audience,
			claims.ExpiresAt.Time,
			claims.NotBefore.Time,
			claims.IssuedAt.Time,
			claims.ID))
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

func NewTokenFromString(tokenString string) (*session.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_PUBLIC_KEY")), nil
	})

	if err != nil {
		return nil, errors.Join(err, errors.New("failed parsing token"))
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)

	if !ok {
		return nil, errors.New("failed parsing token")
	}

	return session.NewSessionToken(tokenString, token.Valid, session.NewClaims(
		claims.Issuer,
		claims.Subject,
		claims.Audience,
		claims.ExpiresAt.Time,
		claims.NotBefore.Time,
		claims.IssuedAt.Time,
		claims.ID,
	))

}
