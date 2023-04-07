package session

import (
	"errors"
	"time"
)

type Token struct {
	Raw    string
	Valid  bool
	Claims *Claims
}

type Claims struct {

	// the `iss` (Issuer) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.1
	Issuer string

	// the `sub` (Subject) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.2
	Subject string

	// the `aud` (Audience) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.3
	Audience []string

	// the `exp` (Expiration Time) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.4
	ExpiresAt time.Time

	// the `nbf` (Not Before) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.5
	NotBefore time.Time

	// the `iat` (Issued At) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.6
	IssuedAt time.Time
}

func NewSessionToken(raw string, valid bool, claims *Claims) (*Token, error) {
	if raw == "" {
		return nil, errors.New("raw is empty")
	}

	if claims == nil {
		return nil, errors.New("claims is empty")
	}

	return &Token{
		Raw:    raw,
		Valid:  valid,
		Claims: claims,
	}, nil
}

func NewClaims(
	issuer string,
	subject string,
	expiresAt time.Time,
	notBefore time.Time,
	issuedAt time.Time) *Claims {
	return &Claims{
		Issuer:    issuer,
		Subject:   subject,
		ExpiresAt: expiresAt,
		NotBefore: notBefore,
		IssuedAt:  issuedAt,
	}
}
