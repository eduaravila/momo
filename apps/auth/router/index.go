package router

import (
	"net/http"

	"github.com/eduaravila/momo/apps/auth/config"
	"github.com/eduaravila/momo/apps/auth/router/oauth"
	"github.com/eduaravila/momo/packages/router"
)

func Routes(env *config.Env) *http.ServeMux {
	mux := router.NewHandler(http.NewServeMux())

	twitchHandler := oauth.NewTwitchHandler(env)

	mux.Get("/oauth/twitch/callback", MakeHTTPHandler(twitchHandler.Callback))
	return mux.GetServeMux()
}

type HttpHandlerFunction func(w http.ResponseWriter, r *http.Request) error

func MakeHTTPHandler(fn HttpHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
