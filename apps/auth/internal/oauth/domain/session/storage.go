package session

import (
	"context"
)

type Storage interface {
	CreateSession(context.Context, Session) (Session, error)
	CreateUserAccount(context.Context, OIDCClaims, OAuthToken) (*UserAccount, error)
}
