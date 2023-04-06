package v1

import (
	"net/http"
	"os"
	"strings"

	"github.com/eduaravila/momo/apps/auth/internal/app"
	"github.com/eduaravila/momo/apps/auth/internal/app/command"
	"github.com/eduaravila/momo/apps/auth/internal/app/query"
	"github.com/eduaravila/momo/apps/auth/internal/domain/session"
	"github.com/eduaravila/momo/apps/auth/internal/port"
	"github.com/google/uuid"
)

type requestIDKey string

type HTTPServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) *HTTPServer {
	return &HTTPServer{app}
}

func (h HTTPServer) OauthTwitchCallback(w http.ResponseWriter, r *http.Request, params port.OauthTwitchCallbackParams) {
	sessionUUID := uuid.NewString()
	err := h.app.Commands.AuthenticateWithOIDC.Handle(r.Context(), command.GenerateSession{
		Code:        params.Code,
		Scope:       strings.Split(params.Scope, " "),
		SessionUUID: sessionUUID,
		AccountUUID: uuid.NewString(),
		UserUUID:    uuid.NewString(),
		Metadata: session.ClientMetadata{
			UserAgent: r.UserAgent(),
			IPAddress: r.RemoteAddr,
		},
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, err := h.app.Queries.SessionWithID.Handle(r.Context(), query.SessionWithID{SessionID: sessionUUID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Name:     "session",
		Value:    session.SessionToken.Raw,
		Path:     "/",
	})

	http.Redirect(w, r, os.Getenv("DASHBOARD_APP_URL"), http.StatusFound)

}
