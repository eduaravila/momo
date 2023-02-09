package v1

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/eduaravila/momo/apps/auth/config"
	"github.com/eduaravila/momo/apps/auth/storage"
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
			r.Context()
		}
	})
}

func TwitchCallback(s *TwitchHandler) HTTPWithError {

	return func(w http.ResponseWriter, r *http.Request) error {
		queryparams := r.URL.Query()
		code := queryparams.Get("code")

		token, err := s.LogIn(code, Metadata{r.UserAgent(), r.RemoteAddr})
		if err != nil {
			return err
		}
		http.SetCookie(w, &http.Cookie{
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
			Name:     "session",
			Value:    token,
			Path:     "/",
		})

		http.Redirect(w, r, os.Getenv("DASHBOARD_APP_URL"), http.StatusFound)
		return nil
	}
}

func Handler(env *config.Env) http.Handler {
	router := router.NewHandler(http.NewServeMux())

	twitchHandler := NewTwitchHandler(storage.NewUserAccountSessionStorage(context.Background(), env.Queries), svc.NewTwitchAPI())
	router.Get("/oauth/twitch/callback", withError(TwitchCallback(twitchHandler)))

	return withCors(router.GetServeMux())
}
