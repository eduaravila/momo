package main

import (
	"log"
	"net/http"

	"github.com/eduaravila/momo/apps/auth/config"
	"github.com/eduaravila/momo/apps/auth/router"
	"github.com/eduaravila/momo/packages/db/queries"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Access-Control-Allow-Origin", "http://localhost")
		r.Header.Set("Access-Control-Allow-Credentials", "true")
		r.Header.Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		next.ServeHTTP(w, r)

	})
}

func main() {

	db, err := config.InitPostgresDB()

	if err != nil {
		log.Fatal(err)
	}
	env := &config.Env{
		Db: db, Queries: queries.New(db),
	}

	mux := http.NewServeMux()

	mux.Handle("/", CorsMiddleware(router.Routes(env)))
	http.ListenAndServe(":8080", mux)
}
