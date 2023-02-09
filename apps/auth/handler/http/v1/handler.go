package v1

import (
	"net/http"
	"os"
	"strings"

	"github.com/eduaravila/momo/apps/auth/service"
	"github.com/eduaravila/momo/apps/auth/storage"
	"github.com/eduaravila/momo/apps/auth/svc"
	"github.com/eduaravila/momo/apps/auth/types"
	"github.com/eduaravila/momo/packages/db/queries"
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

func TwitchCallback(q *queries.Queries, twitchAPI *svc.TwitchAPI) HTTPWithError {

	return func(w http.ResponseWriter, r *http.Request) error {
		s := service.NewTwitchHandler(storage.NewStorage(r.Context(), q), twitchAPI)

		queryparams := r.URL.Query()
		code := queryparams.Get("code")

		token, err := s.LogIn(code, types.Metadata{UserAgent: r.UserAgent(), IPAddress: r.RemoteAddr})
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

func Handler(q *queries.Queries, twitchAPI *svc.TwitchAPI) http.Handler {
	router := router.NewHandler(http.NewServeMux())

	router.Get("/oauth/twitch/callback", withError(TwitchCallback(q, twitchAPI)))

	return withCors(router.GetServeMux())
}
