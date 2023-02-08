package v1

import (
	"net/http"
	"strings"

	"github.com/eduaravila/momo/apps/auth/config"
	"github.com/eduaravila/momo/apps/auth/svc"
	"github.com/eduaravila/momo/packages/router"
)

type HTTPWithError func(w http.ResponseWriter, r *http.Request) error

func withCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var accessControlAllowHeaders strings.Builder
		accessControlAllowHeaders.WriteString("Accept,")
		accessControlAllowHeaders.WriteString("Content-Type,")
		accessControlAllowHeaders.WriteString("Content-Length,")
		accessControlAllowHeaders.WriteString("Accept-Encoding,")
		accessControlAllowHeaders.WriteString("X-CSRF-Token,")
		accessControlAllowHeaders.WriteString("Authorization")

		r.Header = map[string][]string{
			"Access-Control-Allow-Origin":      {"http://localhost"},
			"Access-Control-Allow-Credentials": {"true"},
			"Access-Control-Allow-Headers":     {accessControlAllowHeaders.String()},
		}
		next.ServeHTTP(w, r)

	})
}

func withError(fn HTTPWithError) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

func Handler(env *config.Env) http.Handler {
	router := router.NewHandler(http.NewServeMux())

	twitchHandler := NewTwitchHandler(env, svc.NewTwitchAPI())
	router.Get("/oauth/twitch/callback", withError(twitchHandler.Callback))

	return withCors(router.GetServeMux())
}
