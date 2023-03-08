package storage

import (
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

func (a *Storage) CreateUserAccount(claims types.OIDCClaims, token types.OAuthToken) (*UserAccount, error) {
	result, err := a.queries.GetAccountAndUserBySub(a.ctx, claims.Sub)

	user := queries.User{
		ID:        result.UserID,
		CreatedAt: result.CreatedAt_2,
		UpdatedAt: result.UpdatedAt_2,
	}

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		user, err = a.queries.CreateUser(a.ctx, uuid.New())
		if err != nil {
			return nil, err
		}
	}

	account, err := a.queries.CreateAccount(a.ctx, queries.CreateAccountParams{
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
