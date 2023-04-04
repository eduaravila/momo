package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/eduaravila/momo/apps/auth/internal/adapter"
	"github.com/eduaravila/momo/apps/auth/internal/domain/session"
	"github.com/eduaravila/momo/packages/postgres/queries"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type OauthPostgresStorage struct {
	queries *queries.Queries
}

func NewSessionPostgresStorage(queries *queries.Queries) *OauthPostgresStorage {
	return &OauthPostgresStorage{queries}
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

func (s OauthPostgresStorage) CreateSession(ctx context.Context, session queries.Session) (queries.Session, error) {
	return s.queries.CreateSession(ctx, queries.CreateSessionParams{
		ID:           uuid.New(),
		ExpiredAt:    session.ExpiredAt,
		UserAgent:    session.UserAgent,
		UserID:       session.UserID,
		SessionToken: session.SessionToken,
		IpAddress:    session.IpAddress,
	})
}

func (s OauthPostgresStorage) CreateUser(ctx context.Context) (queries.User, error) {
	return s.queries.CreateUser(ctx,
		uuid.New(),
	)
}

func (o OauthPostgresStorage) GetAccountAndUserBySub(ctx context.Context, sub string) (queries.GetAccountAndUserBySubRow, error) {
	return o.queries.GetAccountAndUserBySub(ctx, sub)
}

func (o *OauthPostgresStorage) AddAccountWithUser(
	ctx context.Context,
	account *session.Account,
	userUUID string,
) error {
	_, err := o.queries.GetAccountAndUserBySub(ctx, account.Sub)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err == sql.ErrNoRows {
		id, err := uuid.FromBytes([]byte(userUUID))
		if err != nil {
			return errors.Join(err, errors.New("error parsing user id"))
		}

		_, err = o.queries.CreateUser(ctx, id)
		if err != nil {
			return errors.Join(err, errors.New("error creating user"))
		}
	}

	parsedAccountUUID, _ := uuid.FromBytes([]byte(account.ID))
	parsedUserUUID, _ := uuid.FromBytes([]byte(account.ID))
	_, err = o.queries.CreateAccount(ctx, queries.CreateAccountParams{
		ID:               parsedAccountUUID,
		UserID:           parsedUserUUID,
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

func (o *OauthPostgresStorage) AddSession(ctx context.Context, session *session.Session) error {
	id, err := uuid.FromBytes([]byte(session.ID))
	if err != nil {
		return err
	}

	userId, err := uuid.FromBytes([]byte(session.UserID))

	if err != nil {
		return err
	}

	_, err = o.queries.CreateSession(ctx, queries.CreateSessionParams{
		ID:           id,
		UserID:       userId,
		CreatedAt:    time.Now(),
		ExpiredAt:    session.ExpiredAt,
		IpAddress:    session.Metadata.IPAddress,
		UserAgent:    session.Metadata.UserAgent,
		IsValid:      session.IsValid,
		SessionToken: session.SessionToken.Raw,
	})

	return err
}

func (o *OauthPostgresStorage) FindUserFromSub(ctx context.Context, sub string) (*session.User, error) {

	dbUser, err := o.queries.GetUserBySub(ctx, sub)
	if err != nil {
		return nil, err
	}

	return session.UnmarshalUserFromDatabase(dbUser.ID.String(), dbUser.CreatedAt, dbUser.UpdatedAt)
}

func (o *OauthPostgresStorage) GetSession(cxt context.Context, sessionId string) (*session.Session, error) {
	sessionIdString, err := uuid.ParseBytes([]byte(sessionId))

	if err != nil {
		return nil, err
	}

	sessionDb, err := o.queries.GetSession(cxt, sessionIdString)
	if err != nil {
		return nil, errors.Join(err, errors.New("error getting session from db"))
	}

	sessionToken, err := adapter.NewTokenFromString(sessionDb.SessionToken)

	if err != nil {
		return nil, errors.Join(err, errors.New("error getting session from db"))
	}

	return session.UnmarshalSessionFromDb(
		sessionDb.ID.String(),
		sessionDb.UserID.String(),
		sessionDb.CreatedAt,
		sessionDb.ExpiredAt,
		sessionDb.IpAddress,
		sessionDb.UserAgent,
		sessionDb.IsValid,
		sessionToken,
	)
}
