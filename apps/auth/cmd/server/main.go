package main

import (
	"log"
	"net/http"

	v1 "github.com/eduaravila/momo/apps/auth/handler/http/v1"
	"github.com/eduaravila/momo/apps/auth/storage"
	"github.com/eduaravila/momo/apps/auth/svc"
	"github.com/eduaravila/momo/packages/db/queries"
)

func main() {
	db, err := storage.InitPostgresDB()

	if err != nil {
		log.Fatal(err)
	}
	router := v1.Handler(queries.New(db), svc.NewTwitchAPI())
	log.Fatal(http.ListenAndServe(":8080", router))
}
