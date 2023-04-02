package query

import "github.com/eduaravila/momo/apps/auth/internal/domain/session"

type SessionWithId struct {
	SessionId string
}

type sessionWithIdHandler struct {
	store session.Storage
}

type SessionWithIdHandler decorators.QueryHandler[SessionWithId]

func NewSessionWithIdHandler(store session.Storage) SessionWithIdHandler {

}
