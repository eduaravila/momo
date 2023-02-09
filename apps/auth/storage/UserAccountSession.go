package storage

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/eduaravila/momo/apps/auth/types"
	"github.com/eduaravila/momo/packages/db/queries"
	"github.com/google/uuid"
)

type UserAccountSessionStorage struct {
	*AccountStorage
	*UserStorage
	*SessionStorage
}

type UserAccountSession struct {
	Account queries.Account
	User    queries.User
	Session queries.Session
}

func NewUserAccountSessionStorage(context context.Context, db *queries.Queries) *UserAccountSessionStorage {
	return &UserAccountSessionStorage{
		SessionStorage: NewSessionStorage(context, db),
		AccountStorage: NewAccountStorage(context, db),
		UserStorage:    NewUserStorage(context, db),
	}
}

func (o *UserAccountSessionStorage) CreateUserAccount(userinfoRespose *types.UserinfoRespose, tokenRespose *types.TokenResponse) (*UserAccountSession, error) {
	result, err := o.GetAccountAndUserBySub(userinfoRespose.Sub)
	uid := result.UserID
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	user := queries.User{
		ID:        result.UserID,
		CreatedAt: result.CreatedAt_2,
		UpdatedAt: result.UpdatedAt_2,
	}
	if err == sql.ErrNoRows {
		user, err = o.CreateUser()
		if err != nil {
			return nil, err
		}
	}
	acc, err := o.CreateAccount(queries.CreateAccountParams{
		AccessToken:      tokenRespose.AccessToken,
		RefreshToken:     tokenRespose.RefreshToken,
		ExpiredAt:        time.Unix(int64(tokenRespose.ExpiresIn), 0),
		Email:            userinfoRespose.Email,
		Picture:          userinfoRespose.Picture,
		Iss:              userinfoRespose.Iss,
		PreferedUsername: userinfoRespose.PreferedUsername,
		ID:               uuid.New(),
		Scope:            strings.Join(tokenRespose.Scope, " "),
		Sub:              userinfoRespose.Sub,
		UserID:           uid,
	})
	if err != nil {
		return nil, err
	}

	return &UserAccountSession{
		User:    user,
		Account: acc,
	}, nil
}
