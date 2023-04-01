package query

import "github.com/eduaravila/momo/apps/auth/internal/domain/session"

type GetSessionWithID struct {
	store session.Storage
}
