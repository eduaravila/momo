package session

import (
	"context"
)

type Storage interface {
	AddSession(context.Context, *Session) error
	AddAccountWithUser(ctx context.Context, account *Account) error
	FindUserFromSub(context.Context, string) (*User, error)
	GetSession(context.Context, string) (*Session, error)
}
