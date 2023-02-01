package config

import (
	"database/sql"

	"github.com/eduaravila/momo/packages/db/queries"
)

type Env struct {
	Queries *queries.Queries
	Db      *sql.DB
}
