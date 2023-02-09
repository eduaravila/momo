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
	*queries.Queries
	context context.Context
}

func NewStorage(context context.Context, queries *queries.Queries) *Storage {
	return &Storage{queries, context}
}

func InitPostgresDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")))

	if err != nil {
		return nil, err
	}
	return db, db.Ping()

}

func (s Storage) CreateSession(session queries.Session) (queries.Session, error) {
	return s.Queries.CreateSession(s.context, queries.CreateSessionParams{
		ID:           uuid.New(),
		ExpiredAt:    session.ExpiredAt,
		UserAgent:    session.UserAgent,
		UserID:       session.UserID,
		SessionToken: session.SessionToken,
		IpAddress:    session.IpAddress,
	})
}

func (s Storage) CreateUser() (queries.User, error) {
	return s.Queries.CreateUser(s.context,
		uuid.New(),
	)
}

func (s Storage) CreateAccount(account queries.Account) (queries.Account, error) {
	return s.Queries.CreateAccount(s.context,
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
	return s.Queries.GetAccountAndUserBySub(s.context, sub)
}
