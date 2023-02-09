package main

import (
	"log"
	"net/http"

	"github.com/eduaravila/momo/apps/auth/config"
	v1 "github.com/eduaravila/momo/apps/auth/handler/http/v1"
	"github.com/eduaravila/momo/apps/auth/storage"
	"github.com/eduaravila/momo/packages/db/queries"
)

func main() {
	db, err := storage.InitPostgresDB()

	if err != nil {
		log.Fatal(err)
	}

	env := &config.Env{
		Db: db, Queries: queries.New(db),
	}

	router := v1.Handler(env)
	log.Fatal(http.ListenAndServe(":8080", router))
}
