package model

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/eduaravila/momo/packages/db/queries"
	"github.com/google/uuid"
)

type TokenBody struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	GrantType    string `json:"grant_type"`
	RedirectURI  string `json:"redirect_uri"`
}

type TokenResponse struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	ExpiresIn    int      `json:"expires_in"`
	TokenType    string   `json:"token_type"`
	Scope        []string `json:"scope"`
}

type UserinfoRespose struct {
	Aud              string    `json:"aud"`
	Exp              int64     `json:"exp"`
	Iat              int64     `json:"iat"`
	Iss              string    `json:"iss"`
	Sub              string    `json:"sub"`
	Email            string    `json:"email"`
	EmailVerified    bool      `json:"email_verified"`
	Picture          string    `json:"picture"`
	PreferedUsername string    `json:"preferred_username"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type Oidc struct {
	User    queries.User
	Account queries.Account
	Db      *queries.Queries
	Context context.Context
}

type OIDCBuilder interface {
	CreateUserAccount(userinfoRespose UserinfoRespose, tokenRespose TokenResponse) (*UserAccount, error)
}

func NewOIDCBuilder(db *queries.Queries, context context.Context) OIDCBuilder {
	return &Oidc{
		Db:      db,
		Context: context,
	}
}

type UserAccount struct {
	User    queries.User
	Account queries.Account
}

func (o *Oidc) CreateUserAccount(userinfoRespose UserinfoRespose, tokenRespose TokenResponse) (*UserAccount, error) {
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
