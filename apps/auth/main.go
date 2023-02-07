package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/eduaravila/momo/apps/auth/config"
	v1 "github.com/eduaravila/momo/apps/auth/handler/v1"
	"github.com/eduaravila/momo/packages/db/queries"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var accessControlAllowHeaders strings.Builder
		accessControlAllowHeaders.WriteString("Accept,")
		accessControlAllowHeaders.WriteString("Content-Type,")
		accessControlAllowHeaders.WriteString("Content-Length,")
		accessControlAllowHeaders.WriteString("Accept-Encoding,")
		accessControlAllowHeaders.WriteString("X-CSRF-Token,")
		accessControlAllowHeaders.WriteString("Authorization")

		r.Header = map[string][]string{
			"Access-Control-Allow-Origin":      {"http://localhost"},
			"Access-Control-Allow-Credentials": {"true"},
			"Access-Control-Allow-Headers":     {accessControlAllowHeaders.String()},
		}
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

	mux.Handle("/", CorsMiddleware(v1.Routes(env)))
	http.ListenAndServe(":8080", mux)
}
