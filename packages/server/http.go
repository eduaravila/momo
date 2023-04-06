package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
)

func RunHTTPServer(
	prefix,
	addrs string,
	createHandler func() http.Handler) {
	rootRouter := http.NewServeMux()
	rootRouter.Handle(prefix, createHandler())

	if err := http.ListenAndServe(addrs, rootRouter); err != nil {
		log.Fatal(err)
	}
}

type HTTPWithError func(w http.ResponseWriter, r *http.Request) error

type ContextKey string

func withRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-Id")
		r = r.WithContext(context.WithValue(r.Context(), ContextKey("requestId"), requestID))
		os.Getenv("HOSTNAME")
		if requestID == "" {
			requestID = uuid.New().String()
		}

		r.Header.Set("X-Request-Id", requestID)

		next.ServeHTTP(w, r)
	})
}

func withCors(next http.Handler) http.Handler {
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

func withError(fn HTTPWithError) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			log.New(os.Stderr, "ERROR: ", log.LstdFlags|log.Lmsgprefix).Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})
}
