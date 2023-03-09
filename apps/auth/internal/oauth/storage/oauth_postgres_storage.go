package storage

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/eduaravila/momo/packages/db/queries"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type OauthPostgresStorage struct {
	queries *queries.Queries
	ctx     context.Context
}

func NewStorage(ctx context.Context, queries *queries.Queries) *OauthPostgresStorage {
	return &OauthPostgresStorage{queries, ctx}
}

func InitPostgresDB() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))

	if err != nil {
		return nil, err
	}

	return db, db.Ping()
}

func (s OauthPostgresStorage) CreateSession(session queries.Session) (queries.Session, error) {
	return s.queries.CreateSession(s.ctx, queries.CreateSessionParams{
		ID:           uuid.New(),
		ExpiredAt:    session.ExpiredAt,
		UserAgent:    session.UserAgent,
		UserID:       session.UserID,
		SessionToken: session.SessionToken,
		IpAddress:    session.IpAddress,
	})
}

func (s OauthPostgresStorage) CreateUser() (queries.User, error) {
	return s.queries.CreateUser(s.ctx,
		uuid.New(),
	)
}

func (s OauthPostgresStorage) CreateAccount(account queries.Account) (queries.Account, error) {
	return s.queries.CreateAccount(s.ctx,
		queries.CreateAccountParams{
			ID:               uuid.New(),
			Sub:              account.Sub,
			Email:            account.Email,
			UserID:           account.UserID,
			Picture:          account.Picture,
			Iss:              account.Iss,
			Scope:            account.Scope,
			ExpiredAt:        account.ExpiredAt,
			PreferedUsername: account.PreferedUsername,
			AccessToken:      account.AccessToken,
			RefreshToken:     account.RefreshToken,
		},
	)
}

func (s OauthPostgresStorage) GetAccountAndUserBySub(sub string) (queries.GetAccountAndUserBySubRow, error) {
	return s.queries.GetAccountAndUserBySub(s.ctx, sub)
}

func (a *OauthPostgresStorage) CreateUserAccount(claims types.OIDCClaims, token types.OAuthToken) (*UserAccount, error) {
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
