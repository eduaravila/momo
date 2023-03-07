package v1

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/eduaravila/momo/apps/auth/internal/adapter"
	"github.com/eduaravila/momo/apps/auth/internal/oidc"
	"github.com/eduaravila/momo/apps/auth/internal/storage"
	"github.com/eduaravila/momo/apps/auth/internal/types"
	"github.com/eduaravila/momo/packages/db/queries"
	"github.com/eduaravila/momo/packages/router"
	"github.com/google/uuid"
)

type requestIDKey string

type HTTPWithError func(w http.ResponseWriter, r *http.Request) error

func withRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-Id")
		r = r.WithContext(context.WithValue(r.Context(), requestIDKey("requestId"), requestID))
		os.Getenv("HOSTNAME")
		if requestID == "" {
			requestID = uuid.New().String()
		}

		r.Header.Set("X-Request-Id", requestID)

		next.ServeHTTP(w, r)
	})
}

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
			log.New(os.Stderr, "ERROR: ", log.LstdFlags|log.Lmsgprefix).Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})
}

func TwitchLogIn(q *queries.Queries, twitchAPI *adapter.TwitchAPI) HTTPWithError {
	return func(w http.ResponseWriter, r *http.Request) error {
		s := oidc.NewAuthService(storage.NewStorage(r.Context(), q), adapter.NewTwitchAPI())

		queryparams := r.URL.Query()
		code := queryparams.Get("code")

		session, err := s.LogIn(code, types.Metadata{UserAgent: r.UserAgent(), IPAddress: r.RemoteAddr})

		if err != nil {
			return err
		}

		http.SetCookie(w, &http.Cookie{
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
			Name:     "session",
			Value:    session.SessionToken,
			Path:     "/",
		})

		http.Redirect(w, r, os.Getenv("DASHBOARD_APP_URL"), http.StatusFound)

		return nil
	}
}

func Handler(q *queries.Queries, twitchAPI *adapter.TwitchAPI) http.Handler {
	router := router.NewHandler(http.NewServeMux())
	router.Get("/oauth/twitch/callback", withRequestID(withError(TwitchLogIn(q, twitchAPI))))

	return withCors(router.GetServeMux())
}
