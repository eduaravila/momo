package router

import (
	"net/http"

	"github.com/eduaravila/momo/apps/auth/config"
	"github.com/eduaravila/momo/apps/auth/router/oauth"
	"github.com/eduaravila/momo/packages/router"
)

func Routes(env *config.Env) *http.ServeMux {
	mux := router.NewHandler(http.NewServeMux())

	twitchRouter := oauth.NewTwitchRouter(env)

	mux.Get("/oauth/twitch/callback", twitchRouter.GetToken)
	return mux.GetServeMux()
}
