package router

import (
	"net/http"
	"os"

	"github.com/eduaravila/momo/apps/auth/api"
	"github.com/eduaravila/momo/apps/auth/config"
	"github.com/eduaravila/momo/apps/auth/router/oauth"
	"github.com/eduaravila/momo/packages/router"
)

func Routes(env *config.Env) *http.ServeMux {
	mux := router.NewHandler(http.NewServeMux())

	twitchHandler := oauth.NewTwitchHandler(env, api.NewTwitchAPI(&api.Config{BaseURL: os.Getenv("TWITCH_OAUTH2_URL")}))

	mux.Get("/oauth/twitch/callback", MakeHTTPHandler(twitchHandler.Callback))
	return mux.GetServeMux()
}

type HTTPWithError func(w http.ResponseWriter, r *http.Request) error

func MakeHTTPHandler(fn HTTPWithError) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
