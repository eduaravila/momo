package oauth

import (
	"context"
	"net/http"

	"github.com/eduaravila/momo/apps/auth/api"
	"github.com/eduaravila/momo/apps/auth/config"
	"github.com/eduaravila/momo/apps/auth/factory"
	"github.com/eduaravila/momo/apps/auth/model"
	"github.com/eduaravila/momo/apps/auth/url"
	"github.com/eduaravila/momo/packages/db/queries"
	"github.com/google/uuid"
)

type TwitchHandler struct {
	env *config.Env
}

func NewTwitchHandler(env *config.Env) *TwitchHandler {
	return &TwitchHandler{env: env}
}

func (t *TwitchHandler) Callback(w http.ResponseWriter, r *http.Request) error {
	queryparams := r.URL.Query()
	code := queryparams.Get("code")

	oidcToken, err := api.GetToken(code)
	if err != nil {
		return err
	}
	userInfo, err := api.GetUserinfo(oidcToken)
	if err != nil {
		return err
	}

	ua, err := model.NewOIDCBuilder(t.env.Queries, r.Context()).CreateUserAccount(*userInfo, *oidcToken)

	if err != nil {
		return err
	}

	token := factory.NewSessionToken(ua.User.ID.String())
	tokenString, err := token.Sign()
	if err != nil {
		return err
	}

	t.env.Queries.CreateSession(context.Background(), queries.CreateSessionParams{
		ID:           uuid.New(),
		ExpiredAt:    token.Claims().ExpiresAt.Time,
		UserAgent:    r.UserAgent(),
		UserID:       ua.User.ID,
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

	http.Redirect(w, r, url.DASHBOARD_APP_URL, http.StatusFound)
	return nil
}
