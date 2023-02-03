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
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	ExpiresIn    int      `json:"expires_in"`
	TokenType    string   `json:"token_type"`
	Scope        []string `json:"scope"`
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
		ClientID:     config.TWITCH_APPLICATION_CLIEND_ID,
		ClientSecret: config.TWITCH_APPLICATION_CLIENT_SECRET,
		Code:         code,
		GrantType:    "authorization_code",
		RedirectURI:  url.DASHBOARD_APP_URL,
	}

	// struct to io.Reader
	jsonBody, err := json.Marshal(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	res, err := Post(PostParams{
		Url:  url.TWITCH_OAUTH2_TOKEN,
		Body: bytes.NewReader(jsonBody),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}
	if res.StatusCode != http.StatusOK {
		http.Redirect(w, r, url.DASHBOARD_APP_URL, http.StatusUnauthorized)
	}

	var tokenRespose TokenResponse
	err = json.NewDecoder(res.Body).Decode(&tokenRespose)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	userInfo, err := Get(PostParams{
		Url: url.TWITCH_OAUTH2_USERINFO,
		Headers: [][]string{
			{"Authorization", "Bearer " + tokenRespose.AccessToken},
		},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, err = io.ReadAll(userInfo.Body)

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

type PostParams struct {
	Url     string
	Body    io.Reader
	Headers [][]string
}

func MakeRequest(method string, params PostParams) (*http.Response, error) {
	request, err := http.NewRequest(method, params.Url, params.Body)
	if err != nil {
		return nil, err
	}
	for _, header := range params.Headers {
		request.Header.Set(header[0], header[1])
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(request)
}

// make a post request to a generic url with a body
func Post(params PostParams) (*http.Response, error) {
	return MakeRequest("POST", params)
}

// make a post request to a generic url with a body
func Get(params PostParams) (*http.Response, error) {
	return MakeRequest("GET", params)
}
