package storage

import (
	"context"

	"github.com/eduaravila/momo/packages/db/queries"
	"github.com/google/uuid"
)

type UserStorage struct {
	queries *queries.Queries
	context context.Context
}

func NewUserStorage(context context.Context, queries *queries.Queries) *UserStorage {

	return &UserStorage{queries, context}
}

func (s *UserStorage) CreateUser() (queries.User, error) {
	uid := uuid.New()
	return s.queries.CreateUser(s.context, uid)
}
