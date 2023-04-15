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
		Subject:   id,
		ExpiresAt: jwt.NewNumericDate(exp),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    os.Getenv("JWT_ISSUER"),
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
			claims.ExpiresAt.Time,
			claims.NotBefore.Time,
			claims.IssuedAt.Time))
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

func (t *SessionToken) VerifyToken(ctx context.Context, tokenString string) (*session.Token, error) {
	decodedKey, err := jwt.ParseECPublicKeyFromPEM([]byte(os.Getenv("JWT_PUBLIC_KEY")))
	if err != nil {
		return nil, err
	}

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return decodedKey, nil
	})

	if err != nil {
		return nil, errors.Join(err, errors.New("failed parsing token"))
	}
	claims := parsedToken.Claims.(jwt.RegisteredClaims)
	claims.Valid()
	return &session.Token{
		Valid: parsedToken.Valid,
		Claims: session.NewClaims(
			claims.Issuer,
			claims.Subject,
			claims.ExpiresAt.Time,
			claims.NotBefore.Time,
			claims.IssuedAt.Time,
		)}, nil
}

func NewTokenFromString(tokenString string) (*session.Token, error) {
	decodedPublicKey, err := jwt.ParseECPublicKeyFromPEM([]byte(os.Getenv("JWT_PUBLIC_KEY")))
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return decodedPublicKey, nil
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
		claims.ExpiresAt.Time,
		claims.NotBefore.Time,
		claims.IssuedAt.Time,
	))

}
