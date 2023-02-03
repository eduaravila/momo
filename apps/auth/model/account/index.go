package account

import (
	"context"

	"github.com/eduaravila/momo/packages/db/queries"
	"github.com/google/uuid"
)

func Create(queries *queries.Queries, params queries.CreateAccountParams) (uuid.UUID, error) {
	return queries.CreateAccount(context.Background(), params)
}
