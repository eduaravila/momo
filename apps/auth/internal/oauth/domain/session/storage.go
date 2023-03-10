package session

import (
	"context"
)

type Storage interface {
	AddSession(context.Context, *Session) error
	AddAccountWithUser(context.Context, OIDCClaims, *Account) error
}
