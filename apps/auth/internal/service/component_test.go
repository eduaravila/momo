package service

import (
	"net/http"
	"os"
	"testing"

	"github.com/eduaravila/momo/apps/auth/internal/ports"
	v1 "github.com/eduaravila/momo/apps/auth/internal/ports/v1"
	"github.com/eduaravila/momo/packages/server"
	"github.com/eduaravila/momo/packages/test"
)

func startService() bool {
	port := os.Getenv("AUTH_PORT")
	if port == "" {
		port = "8080"
	}

	app := NewApplication()
	authAddress := os.Getenv("AUTH_HTTP_ADDR")

	go server.RunHTTPServer("/api", ":"+port, func() http.Handler {
		return ports.Handler(v1.NewHTTPServer(app))
	})

	ok := test.WaitFor(authAddress)

	if !ok {
		return false
	}

	return true
}

func TestMain(m *testing.M) {
	if !startService() {
		os.Exit(1)
	}

	os.Exit(m.Run())
}
