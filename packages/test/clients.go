package test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/eduaravila/momo/packages/client/auth"
	"github.com/stretchr/testify/require"
)

type AuthHTTPClient struct {
	client *auth.ClientWithResponses
}

func NewAuthHTTPClient(t *testing.T) AuthHTTPClient {
	url := os.Getenv("AUTH_HTTP_ADDR")

	c, err := auth.NewClientWithResponses(url)
	require.NoError(t, err)

	return AuthHTTPClient{
		client: c,
	}
}

func (a *AuthHTTPClient) ShouldAuthenticateWithTwitch(
	t *testing.T,
	code,
	scope string,
	expectedSessionToken string,
) error {
	res, err := a.client.OauthTwitchCallbackWithResponse(context.Background(), &auth.OauthTwitchCallbackParams{
		Code:  code,
		Scope: scope,
	})

	require.NoError(t, err)
	require.Equal(t, http.StatusFound, res.StatusCode)
	sessionCookie := res.HTTPResponse.Cookies()[0]

	require.Equal(t, sessionCookie.Value, expectedSessionToken)

	return err
}
