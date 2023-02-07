package v1

import (
	"context"
	"net/http"
	"os"

	"github.com/eduaravila/momo/apps/auth/config"
	"github.com/eduaravila/momo/apps/auth/factory"
	"github.com/eduaravila/momo/apps/auth/storage"
	"github.com/eduaravila/momo/apps/auth/svc"
	"github.com/eduaravila/momo/packages/db/queries"
	"github.com/google/uuid"
)

type TwitchHandler struct {
	env *config.Env
	api *svc.TwitchAPI
}

func NewTwitchHandler(env *config.Env, api *svc.TwitchAPI) *TwitchHandler {
	return &TwitchHandler{env: env, api: api}
}

func (t *TwitchHandler) Callback(w http.ResponseWriter, r *http.Request) error {

	queryparams := r.URL.Query()
	code := queryparams.Get("code")

	oidcToken, err := t.api.GetToken(code)

	if err != nil {
		return err
	}
	userInfo, err := t.api.GetOidcUserInfo(oidcToken)
	if err != nil {
		return err
	}

	userAndAccount, err := storage.NewOIDCBuilder(r.Context(), t.env.Queries).CreateUserAccount(*userInfo, *oidcToken)

	if err != nil {
		return err
	}

	token := factory.NewSessionToken(userAndAccount.User.ID.String())
	tokenString, err := token.Sign()
	if err != nil {
		return err
	}

	t.env.Queries.CreateSession(context.Background(), queries.CreateSessionParams{
		ID:           uuid.New(),
		ExpiredAt:    token.Claims().ExpiresAt.Time,
		UserAgent:    r.UserAgent(),
		UserID:       userAndAccount.User.ID,
		SessionToken: tokenString,
		IpAddress:    r.RemoteAddr,
	})

	http.SetCookie(w, &http.Cookie{
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Name:     "session",
		Value:    tokenString,
		Path:     "/",
	})

	http.Redirect(w, r, os.Getenv("DASHBOARD_APP_URL"), http.StatusFound)
	return nil
}
