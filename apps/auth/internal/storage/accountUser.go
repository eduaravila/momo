package storage

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/eduaravila/momo/apps/auth/internal/types"
	"github.com/eduaravila/momo/packages/db/queries"
	"github.com/google/uuid"
)

type UserAccount struct {
	User    queries.User
	Account queries.Account
}

type AccountUserStorage struct {
	queries *queries.Queries
	context context.Context
}

func NewAccountUserStorage(context context.Context, queries *queries.Queries) *AccountUserStorage {
	return &AccountUserStorage{queries, context}
}

func (a *Storage) CreateUserAccount(claims types.OIDCClaims, token types.OAuthToken) (*UserAccount, error) {
	result, err := a.queries.GetAccountAndUserBySub(a.context, claims.Sub)
	uid := result.UserID
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	var user queries.User
	if err == sql.ErrNoRows {
		user, err = a.queries.CreateUser(a.context, uuid.New())
		if err != nil {
			return nil, err
		}
	}

	account, err := a.queries.CreateAccount(a.context, queries.CreateAccountParams{
		AccessToken:      token.AccessToken,
		RefreshToken:     token.RefreshToken,
		ExpiredAt:        time.Unix(int64(token.ExpiresIn), 0),
		Email:            claims.Email,
		Picture:          claims.Picture,
		Iss:              claims.Iss,
		PreferedUsername: claims.PreferedUsername,
		ID:               uuid.New(),
		Scope:            strings.Join(token.Scope, " "),
		Sub:              claims.Sub,
		UserID:           uid,
	})
	if err != nil {
		return nil, err
	}

	return &UserAccount{user, account}, nil
}
