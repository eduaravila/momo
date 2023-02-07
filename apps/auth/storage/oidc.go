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

type Oidc struct {
	User    queries.User
	Account queries.Account
	Db      *queries.Queries
	Context context.Context
}

type OIDCBuilder interface {
	CreateUserAccount(userinfoRespose types.UserinfoRespose, tokenRespose types.TokenResponse) (*UserAccount, error)
}

func NewOIDCBuilder(context context.Context, db *queries.Queries) OIDCBuilder {
	return &Oidc{
		Db:      db,
		Context: context,
	}
}

type UserAccount struct {
	User    queries.User
	Account queries.Account
}

func (o *Oidc) CreateUserAccount(userinfoRespose types.UserinfoRespose, tokenRespose types.TokenResponse) (*UserAccount, error) {
	result, err := o.Db.GetAccountAndUserBySub(o.Context, userinfoRespose.Sub)
	uid := result.UserID
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	o.User = queries.User{
		ID:        result.UserID,
		CreatedAt: result.CreatedAt_2,
		UpdatedAt: result.UpdatedAt_2,
	}
	if err == sql.ErrNoRows {
		uid = uuid.New()
		user, err := o.Db.CreateUser(o.Context, uid)
		if err != nil {
			return nil, err
		}
		o.User = user
	}
	acc, err := o.Db.CreateAccount(o.Context, queries.CreateAccountParams{
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

	return &UserAccount{
		User:    o.User,
		Account: acc,
	}, nil
}
