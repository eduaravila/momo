package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/eduaravila/momo/apps/auth/internal/oauth/domain/session"
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

func (s OauthPostgresStorage) GetAccountAndUserBySub(sub string) (queries.GetAccountAndUserBySubRow, error) {
	return s.queries.GetAccountAndUserBySub(s.ctx, sub)
}

func (a *OauthPostgresStorage) AddAccountWithUser(ctx context.Context, account session.Account) error {
	_, err := a.queries.GetAccountAndUserBySub(a.ctx, account.Sub)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err == sql.ErrNoRows {
		_, err = a.queries.CreateUser(a.ctx, uuid.New())
		if err != nil {
			return errors.Join(err, errors.New("error creating user"))
		}
	}

	accountUUID, _ := uuid.FromBytes([]byte(account.ID))
	userUUID, _ := uuid.FromBytes([]byte(account.ID))
	_, err = a.queries.CreateAccount(ctx, queries.CreateAccountParams{
		ID:               accountUUID,
		UserID:           userUUID,
		Sub:              account.Sub,
		Email:            account.Email,
		PreferedUsername: account.PreferedUsername,
		UpdatedAt:        time.Now(),
		CreatedAt:        time.Now(),
		Picture:          account.Picture,
		AccessToken:      account.AccessToken,
		RefreshToken:     account.RefreshToken,
		Iss:              account.Iss,
		ExpiredAt:        account.ExpiredAt,
		Scope:            strings.Join(account.Scope, " "),
	})

	if err != nil {
		return errors.Join(err, errors.New("error creating account"))
	}

	return nil
}
