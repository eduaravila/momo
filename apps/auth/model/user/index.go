package user

import (
	"context"

	"github.com/eduaravila/momo/packages/db/queries"
	"github.com/google/uuid"
)

func Create(queries *queries.Queries) (uuid.UUID, error) {
	id := uuid.New()
	return id, queries.CreateUser(context.Background(), id)
}
