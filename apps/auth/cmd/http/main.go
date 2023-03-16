package main

import (
	"net/http"
	"os"

	"github.com/eduaravila/momo/packages/server"
)

func main() {
	port := os.Getenv("AUTH_PORT")
	server.RunHTTPServer("/api", port, func() http.Handler {
		return http.NewServeMux()
	})
}
