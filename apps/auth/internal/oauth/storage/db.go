package storage

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/eduaravila/momo/packages/db/queries"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Storage struct {
	queries *queries.Queries
	ctx     context.Context
}

func NewStorage(ctx context.Context, queries *queries.Queries) *Storage {
	return &Storage{queries, ctx}
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

func (s Storage) CreateSession(session queries.Session) (queries.Session, error) {
	return s.queries.CreateSession(s.ctx, queries.CreateSessionParams{
		ID:           uuid.New(),
		ExpiredAt:    session.ExpiredAt,
		UserAgent:    session.UserAgent,
		UserID:       session.UserID,
		SessionToken: session.SessionToken,
		IpAddress:    session.IpAddress,
	})
}

func (s Storage) CreateUser() (queries.User, error) {
	return s.queries.CreateUser(s.ctx,
		uuid.New(),
	)
}

func (s Storage) CreateAccount(account queries.Account) (queries.Account, error) {
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

func (s Storage) GetAccountAndUserBySub(sub string) (queries.GetAccountAndUserBySubRow, error) {
	return s.queries.GetAccountAndUserBySub(s.ctx, sub)
}
