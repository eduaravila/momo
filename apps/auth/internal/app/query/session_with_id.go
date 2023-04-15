package query

import (
	"context"

	"github.com/eduaravila/momo/apps/auth/internal/domain/session"
	"github.com/eduaravila/momo/packages/decorators"
)

type SessionWithID struct {
	SessionID string
}

type sessionWithIDHandler struct {
	store session.Storage
}

type SessionWithIDHandler decorators.QueryHandler[SessionWithID, *session.Session]

func NewSessionWithIDHandler(store session.Storage) SessionWithIDHandler {
	return &sessionWithIDHandler{store}
}

func (s *sessionWithIDHandler) Handle(
	cxt context.Context,
	query SessionWithID,
) (*session.Session, error) {
	return s.store.GetSession(cxt, query.SessionID)
}
