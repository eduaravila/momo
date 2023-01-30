package twitch

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	urls "github.com/eduaravila/momo/apps/auth/constant"
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

func GetToken(w http.ResponseWriter, r *http.Request) {
	queryparams := r.URL.Query()
	code := queryparams.Get("code")

	body := TokenBody{
		ClientID:     urls.TWITCH_APPLICATION_CLIEND_ID,
		ClientSecret: urls.TWITCH_APPLICATION_CLIENT_SECRET,
		Code:         code,
		GrantType:    "authorization_code",
		RedirectURI:  urls.DASHBOARD_APP_URL,
	}

	// struct to io.Reader
	jsonBody, err := json.Marshal(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	res, err := Post(urls.TWITCH_OAUTH2_URL, bytes.NewReader(jsonBody))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	if res.StatusCode != http.StatusOK {
		http.Redirect(w, r, urls.DASHBOARD_APP_URL, http.StatusUnauthorized)
	}

	http.Redirect(w, r, urls.DASHBOARD_APP_URL, http.StatusCreated)
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
