package storage

import (
	"context"

	"github.com/eduaravila/momo/packages/postgres/queries"
	"github.com/google/uuid"
)

type SessionStorage struct {
	queries *queries.Queries
	context context.Context
}

func NewSessionStorage(context context.Context, queries *queries.Queries) *SessionStorage {
	return &SessionStorage{queries, context}
}

func (s *SessionStorage) CreateSession(session queries.Session) (queries.Session, error) {
	return s.queries.CreateSession(s.context, queries.CreateSessionParams{
		ID:           uuid.New(),
		ExpiredAt:    session.ExpiredAt,
		UserAgent:    session.UserAgent,
		UserID:       session.UserID,
		SessionToken: session.SessionToken,
		IpAddress:    session.IpAddress,
	})
}
