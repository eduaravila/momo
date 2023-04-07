package main

import (
	"net/http"
	"os"

	"github.com/eduaravila/momo/apps/auth/internal/ports"
	v1 "github.com/eduaravila/momo/apps/auth/internal/ports/v1"
	"github.com/eduaravila/momo/apps/auth/internal/service"
	"github.com/eduaravila/momo/packages/server"
)

func main() {
	port := os.Getenv("AUTH_PORT")
	if port == "" {
		port = "8080"
	}

	app := service.NewApplication()

	server.RunHTTPServer("/api", ":"+port, func() http.Handler {
		return ports.Handler(v1.NewHTTPServer(app))
	})

}
