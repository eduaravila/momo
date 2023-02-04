package model

import (
	"context"

	"github.com/eduaravila/momo/packages/db/queries"
)

func CreateSession(db *queries.Queries, params queries.CreateSessionParams) (queries.Session, error) {
	return db.CreateSession(context.Background(), params)
}
