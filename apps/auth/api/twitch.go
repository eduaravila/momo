package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/eduaravila/momo/apps/auth/config"
	"github.com/eduaravila/momo/apps/auth/model"
	"github.com/eduaravila/momo/apps/auth/url"
	"github.com/eduaravila/momo/apps/auth/utils"
)

type Config interface {
	SetBaseURL(baseurl string)
	GetBaseURL() string
}

type ApiConfig struct {
	BaseURL string
}

func (a *ApiConfig) SetBaseURL(baseurl string) {
	a.BaseURL = baseurl
}

func (a *ApiConfig) GetBaseURL() string {
	return a.BaseURL
}

type TwitchApi struct {
	Config
}

// Paths for the Twitch API
const (
	token_path    = "/oauth2/token"    // POST
	userinfo_path = "/oauth2/userinfo" // GET
)

func (t *TwitchApi) GetToken(code string) (*model.TokenResponse, error) {

	body := model.TokenBody{
		ClientID:     config.TWITCH_APPLICATION_CLIEND_ID,
		ClientSecret: config.TWITCH_APPLICATION_CLIENT_SECRET,
		Code:         code,
		GrantType:    "authorization_code",
		RedirectURI:  url.DASHBOARD_APP_URL,
	}

	// struct to io.Reader
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	res, err := utils.Post(utils.RequestParams{
		Url:  fmt.Sprintf("%s%s", t.Config.GetBaseURL(), token_path),
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

func (t *TwitchApi) GetOidcUserInfo(oidcToken *model.TokenResponse) (*model.UserinfoRespose, error) {

	// get user info
	userInfo, err := utils.Get(utils.RequestParams{
		Url: fmt.Sprintf("%s%s", url.TWITCH_API_URL, userinfo_path),
		Headers: [][]string{
			{"Authorization", "Bearer " + oidcToken.AccessToken},
		},
		Body: nil,
	})

	var userInfoRespose model.UserinfoRespose
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(userInfo.Body).Decode(&userInfoRespose)
	if err != nil {
		return nil, err
	}
	return &userInfoRespose, nil
}

func NewTwitchApi(config Config) *TwitchApi {
	return &TwitchApi{
		Config: config,
	}
}

var TwitchApiWithConfig = NewTwitchApi(&ApiConfig{BaseURL: url.TWITCH_API_URL})
