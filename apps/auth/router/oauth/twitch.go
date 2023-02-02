package oauth

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/eduaravila/momo/apps/auth/config"
	"github.com/eduaravila/momo/apps/auth/model/user"
	"github.com/eduaravila/momo/apps/auth/url"
	"github.com/golang-jwt/jwt/v4"
)

type TokenBody struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	GrantType    string `json:"grant_type"`
	RedirectURI  string `json:"redirect_uri"`
}

type TokenResponse struct {
	access_token  string `json:"access_token"`
	refresh_token string `json:"refresh_token"`
	expires_in    int    `json:"expires_in"`
	token_type    string `json:"token_type"`
	scope         string `json:"scope"`
}

type TwitchHandler struct {
	env *config.Env
}

func NewTwitchHandler(env *config.Env) *TwitchHandler {
	return &TwitchHandler{env: env}
}

func (t *TwitchHandler) GetToken(w http.ResponseWriter, r *http.Request) {
	queryparams := r.URL.Query()
	code := queryparams.Get("code")

	body := TokenBody{
		ClientID:     url.TWITCH_APPLICATION_CLIEND_ID,
		ClientSecret: url.TWITCH_APPLICATION_CLIENT_SECRET,
		Code:         code,
		GrantType:    "authorization_code",
		RedirectURI:  url.DASHBOARD_APP_URL,
	}

	// struct to io.Reader
	jsonBody, err := json.Marshal(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	res, err := Post(url.TWITCH_OAUTH2_URL, bytes.NewReader(jsonBody))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	if res.StatusCode != http.StatusOK {
		http.Redirect(w, r, url.DASHBOARD_APP_URL, http.StatusUnauthorized)
	}

	uid, err := user.Create(t.env.Queries)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	exp := time.Now().Add(1 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"uid": uid.String(),
		"exp": exp.Unix(),
		"iat": time.Now().Unix(),
	})
	key, err := ioutil.ReadFile(".keys/private.gem")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	decodedKey, err := jwt.ParseECPrivateKeyFromPEM(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	session, err := token.SignedString(decodedKey)
	// session cookie
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.SetCookie(w, &http.Cookie{
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Name:     "session",
		Value:    session,
		Path:     "/",
	})

	http.Redirect(w, r, url.DASHBOARD_APP_URL, http.StatusFound)

}

// make a post request to a generic url with a body
func Post(url string, body io.Reader) (*http.Response, error) {
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	return client.Do(request)
}
