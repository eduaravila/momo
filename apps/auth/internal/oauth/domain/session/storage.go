package session

import (
	"context"
)

type Storage interface {
	AddSession(context.Context, *Session) error
	AddAccountWithUser(context.Context, *Account) error
	FindUserFromSub(context.Context, string) (*User, error)
}
