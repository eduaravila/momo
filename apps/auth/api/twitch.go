package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/eduaravila/momo/apps/auth/config"
	"github.com/eduaravila/momo/apps/auth/model"
	"github.com/eduaravila/momo/apps/auth/utils"
)

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

func (t *TwitchAPI) GetToken(code string) (*model.TokenResponse, error) {
	body := model.TokenBody{
		ClientID:     config.TWITCH_APPLICATION_CLIEND_ID,
		ClientSecret: config.TWITCH_APPLICATION_CLIENT_SECRET,
		Code:         code,
		GrantType:    "authorization_code",
		RedirectURI:  os.Getenv("DASHBOARD_APP_URL"),
	}

	// struct to io.Reader
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	res, err := utils.Post(utils.RequestParams{
		Url:  fmt.Sprintf("%s%s", t.IConfig.GetBaseURL(), tokenPath),
		Body: bytes.NewReader(jsonBody),
	})

	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("gettoken: error getting token")
	}

	var tokenRespose model.TokenResponse
	err = json.NewDecoder(res.Body).Decode(&tokenRespose)
	if err != nil {
		return nil, err
	}
	return &tokenRespose, nil
}

func (t *TwitchAPI) GetOidcUserInfo(oidcToken *model.TokenResponse) (*model.UserinfoRespose, error) {
	// get user info
	userInfo, err := utils.Get(utils.RequestParams{
		Url: fmt.Sprintf("%s%s", os.Getenv("TWITCH_API_URL"), userInfoPath),
		Headers: [][]string{
			{"Authorization", "Bearer " + oidcToken.AccessToken},
		},
		Body: nil,
	})
	if err != nil {
		return nil, err
	}

	var userInfoRespose model.UserinfoRespose

	if err = json.NewDecoder(userInfo.Body).Decode(&userInfoRespose); err != nil {
		return nil, err
	}

	return &userInfoRespose, nil
}

func NewTwitchAPI() *TwitchAPI {
	return &TwitchAPI{
		IConfig: &Config{BaseURL: os.Getenv("TWITCH_OAUTH2_URL")},
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
