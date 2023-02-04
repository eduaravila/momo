package config

import (
	"context"
	"database/sql"
	"os"

	"github.com/eduaravila/momo/packages/db/queries"
)

type Env struct {
	Queries *queries.Queries
	Db      *sql.DB
}

var TWITCH_APPLICATION_CLIEND_ID string = os.Getenv("TWITCH_APPLICATION_CLIEND_ID")

var TWITCH_APPLICATION_CLIENT_SECRET string = os.Getenv("TWITCH_APPLICATION_CLIENT_SECRET")

var Ctx = context.Background()
