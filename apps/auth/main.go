package main

import (
	"log"
	"net/http"

	"github.com/eduaravila/momo/apps/auth/config"
	"github.com/eduaravila/momo/apps/auth/router"
	"github.com/eduaravila/momo/packages/db/queries"
)

func main() {

	db, err := config.InitPostgresDB()

	if err != nil {
		log.Fatal(err)
	}
	env := &config.Env{
		Db: db, Queries: queries.New(db),
	}

	mux := http.NewServeMux()
	mux.Handle("/", router.Routes(env))
	http.ListenAndServe(":8080", mux)
}
