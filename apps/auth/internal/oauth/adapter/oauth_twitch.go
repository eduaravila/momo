package adapter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/eduaravila/momo/packages/router"
)

type (
	TokenBodyRequest struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		Code         string `json:"code"`
		GrantType    string `json:"grant_type"`
		RedirectURI  string `json:"redirect_uri"`
	}

	OAuthToken struct {
		AccessToken  string   `json:"access_token"`
		RefreshToken string   `json:"refresh_token"`
		ExpiresIn    int      `json:"expires_in"`
		TokenType    string   `json:"token_type"`
		Scope        []string `json:"scope"`
	}
)

type OIDCClaims struct {
	Aud              string    `json:"aud"`
	Exp              int64     `json:"exp"`
	Iat              int64     `json:"iat"`
	Iss              string    `json:"iss"`
	Sub              string    `json:"sub"`
	Email            string    `json:"email"`
	EmailVerified    bool      `json:"email_verified"`
	Picture          string    `json:"picture"`
	PreferedUsername string    `json:"preferred_username"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type Metadata struct {
	UserAgent string
	IPAddress string
}

type IConfig interface {
	SetBaseURL(baseurl string)
	GetBaseURL() string
}

type Config struct {
	BaseURL string
}

func (a *Config) SetBaseURL(baseurl string) {
	a.BaseURL = baseurl
}

func (a *Config) GetBaseURL() string {
	return a.BaseURL
}

type TwitchAPI struct {
	IConfig
}

// Paths for the Twitch API
const (
	tokenPath    = "/oauth2/token"    // POST
	userInfoPath = "/oauth2/userinfo" // GET
)

func (t *TwitchAPI) GetToken(code string) (*OAuthToken, error) {
	body := TokenBodyRequest{
		ClientID:     os.Getenv("TWITCH_APPLICATION_CLIEND_ID"),
		ClientSecret: os.Getenv("TWITCH_APPLICATION_CLIENT_SECRET"),
		Code:         code,
		GrantType:    "authorization_code",
		RedirectURI:  os.Getenv("DASHBOARD_APP_URL"),
	}

	// struct to io.Reader
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	res, err := router.Post(router.RequestParams{
		Url:  fmt.Sprintf("%s%s", t.IConfig.GetBaseURL(), tokenPath),
		Body: bytes.NewReader(jsonBody),
	})

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("gettoken: error getting token")
	}

	var tokenRespose OAuthToken
	err = json.NewDecoder(res.Body).Decode(&tokenRespose)

	if err != nil {
		return nil, err
	}

	return &tokenRespose, nil
}

func (t *TwitchAPI) GetOIDCUserInfo(oidcToken *OAuthToken) (*OIDCClaims, error) {
	// get user info
	userInfo, err := router.Get(router.RequestParams{
		Url: fmt.Sprintf("%s%s", os.Getenv("TWITCH_API_URL"), userInfoPath),
		Headers: [][]string{
			{"Authorization", "Bearer " + oidcToken.AccessToken},
		},
		Body: nil,
	})
	if err != nil {
		return nil, err
	}

	var userInfoRespose OIDCClaims

	if err = json.NewDecoder(userInfo.Body).Decode(&userInfoRespose); err != nil {
		return nil, err
	}

	return &userInfoRespose, nil
}

func NewTwitchAPI() *TwitchAPI {
	return &TwitchAPI{
		IConfig: &Config{BaseURL: os.Getenv("TWITCH_API_URL")},
	}
}

func NewTwitchAPIWithOpts(config IConfig) *TwitchAPI {
	if config != nil {
		return &TwitchAPI{
			IConfig: config,
		}
	}

	return NewTwitchAPI()
}
