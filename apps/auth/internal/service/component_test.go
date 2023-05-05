package service

import (
	"net/http"
	"os"
	"testing"

	"github.com/eduaravila/momo/apps/auth/internal/ports"
	v1 "github.com/eduaravila/momo/apps/auth/internal/ports/v1"
	"github.com/eduaravila/momo/packages/server"
	"github.com/eduaravila/momo/packages/test"
	"github.com/stretchr/testify/require"
)

func startService() bool {
	port := os.Getenv("AUTH_PORT")
	if port == "" {
		port = "8080"
	}

	app := NewTestApplication()
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

func TestAuthenticateWithTwitch(t *testing.T) {
	t.Parallel()
	addr := os.Getenv("AUTH_HTTP_ADDR")
	client := test.NewAuthHTTPClient(t)
	ok := test.WaitFor(addr)
	require.True(t, ok, "auth HTTP time out.")
	// TODO: create mock session token generator, in mock services
	client.ShouldAuthenticateWithTwitch(t, "code", "scope", "session-token")
}

func TestMain(m *testing.M) {
	if !startService() {
		os.Exit(1)
	}

	os.Exit(m.Run())
}
