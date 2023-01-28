package router

import (
	"net/http"

	"github.com/eduaravila/momo/auth/router/twitch"
	"github.com/eduaravila/momo/packages/router"
)

func Routes() *http.ServeMux {
	mux := router.NewHandler(http.NewServeMux())
	mux.Get("/oauth/twitch/callback", twitch.GetToken)
	return mux.GetServeMux()
}
