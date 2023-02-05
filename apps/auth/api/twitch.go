package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/eduaravila/momo/apps/auth/config"
	"github.com/eduaravila/momo/apps/auth/model"
	"github.com/eduaravila/momo/apps/auth/url"
	"github.com/eduaravila/momo/apps/auth/utils"
)

func GetToken(code string) (*model.TokenResponse, error) {

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
		Url:  url.TWITCH_OAUTH2_TOKEN,
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

func GetUserinfo(oidcToken *model.TokenResponse) (*model.UserinfoRespose, error) {

	// get user info
	userInfo, err := utils.Get(utils.RequestParams{
		Url: url.TWITCH_OAUTH2_USERINFO,
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
