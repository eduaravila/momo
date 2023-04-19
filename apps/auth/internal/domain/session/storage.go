package session

import (
	"context"
)

type Storage interface {
	AddSession(context.Context, *Session) error
	GetOrCreateUserFromSub(ctx context.Context, userID string, sub string) (*User, error)
	AddAccountWithUser(ctx context.Context, account *Account, user *User) error
	FindUserFromSub(context.Context, string) (*User, error)
	GetSession(context.Context, string) (*Session, error)
}
