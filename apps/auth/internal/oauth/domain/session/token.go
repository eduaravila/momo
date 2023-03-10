package session

import "time"

type Token struct {
	raw    string
	valid  bool
	claims Claims
}

type Claims struct {

	// the `iss` (Issuer) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.1
	issuer string

	// the `sub` (Subject) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.2
	subject string

	// the `aud` (Audience) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.3
	audience []string

	// the `exp` (Expiration Time) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.4
	expiresAt time.Time

	// the `nbf` (Not Before) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.5
	notBefore time.Time

	// the `iat` (Issued At) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.6
	issuedAt time.Time

	// the `jti` (JWT ID) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.7
	id string
}

func NewSessionToken(raw string, valid bool, claims Claims) Token {
	return Token{
		raw:    raw,
		valid:  valid,
		claims: claims,
	}
}

func NewClaims(issuer string, subject string, audience []string, expiresAt time.Time, notBefore time.Time, issuedAt time.Time, id string) Claims {
	return Claims{
		issuer:    issuer,
		subject:   subject,
		audience:  audience,
		expiresAt: expiresAt,
		notBefore: notBefore,
		issuedAt:  issuedAt,
		id:        id,
	}
}
