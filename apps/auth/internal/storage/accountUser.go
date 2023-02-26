package storage

import (
	"context"
	"database/sql"
	"fmt"
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
	context context.Context
	queries *queries.Queries
}

func NewAccountUserStorage(context context.Context, queries *queries.Queries) *AccountUserStorage {
	return &AccountUserStorage{context, queries}
}

func (a *Storage) CreateUserAccount(claims types.OIDCClaims, token types.OAuthToken) (*UserAccount, error) {
	result, err := a.queries.GetAccountAndUserBySub(a.context, claims.Sub)

	user := queries.User{
		ID:        result.UserID,
		CreatedAt: result.CreatedAt_2,
		UpdatedAt: result.UpdatedAt_2,
	}

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		user, err = a.queries.CreateUser(a.context, uuid.New())
		if err != nil {
			return nil, err
		}
	}

	fmt.Println(strings.Join(token.Scope, " "), time.Now().Add(time.Duration(int64(token.ExpiresIn))*time.Second))

	account, err := a.queries.CreateAccount(a.context, queries.CreateAccountParams{
		ID:               uuid.New(),
		UserID:           user.ID,
		Picture:          claims.Picture,
		Email:            claims.Email,
		PreferedUsername: claims.PreferedUsername,
		AccessToken:      token.AccessToken,
		RefreshToken:     token.RefreshToken,
		Iss:              claims.Iss,
		Sub:              claims.Sub,
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
		ExpiredAt:        time.Now().Add(time.Duration(int64(token.ExpiresIn)) * time.Second),
		Scope:            strings.Join(token.Scope, " "),
	})

	if err != nil {
		return nil, err
	}

	return &UserAccount{user, account}, nil
}
