package main

import (
	"log"
	"net/http"
	"os"

	"github.com/eduaravila/momo/packages/server"
)

func main() {
	port := os.Getenv("AUTH_PORT")
	if port == "" {
		port = ":8080"
	}

	server.RunHTTPServer("/api", port, func() http.Handler {
		log.Default().Println("Listening on port", port)
		return http.NewServeMux()
	})

}
