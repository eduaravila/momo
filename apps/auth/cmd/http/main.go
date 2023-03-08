package main

import (
	"log"
	"net/http"

	"github.com/eduaravila/momo/apps/auth/internal/adapter"
	v1 "github.com/eduaravila/momo/apps/auth/internal/handler/http"
	"github.com/eduaravila/momo/apps/auth/internal/storage"
	"github.com/eduaravila/momo/packages/db/queries"
)

func main() {
	db, err := storage.InitPostgresDB()

	if err != nil {
		log.Fatal(err)
	}

	router := v1.Handler(queries.New(db), adapter.NewTwitchAPI())
	log.Fatal(http.ListenAndServe(":8080", router))
}
