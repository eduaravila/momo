package storage

import (
	"context"

	"github.com/eduaravila/momo/packages/db/queries"
)

type AccountStorage struct {
	queries *queries.Queries
	context context.Context
}

func NewAccountStorage(context context.Context, queries *queries.Queries) *AccountStorage {
	return &AccountStorage{queries, context}
}

func (s *AccountStorage) CreateAccount(account queries.CreateAccountParams) (queries.Account, error) {
	return s.queries.CreateAccount(s.context, account)
}

func (s *AccountStorage) GetAccountAndUserBySub(sub string) (queries.GetAccountAndUserBySubRow, error) {
	return s.queries.GetAccountAndUserBySub(s.context, sub)
}
