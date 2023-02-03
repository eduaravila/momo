package url

import (
	"os"
)

// DASHBOARD_APP_URL app client id
var DASHBOARD_APP_URL string = os.Getenv("DASHBOARD_APP_URL")

// TWITCH_TOKEN_URL app client id
var TWITCH_OAUTH2_URL string = os.Getenv("TWITCH_OAUTH2_URL")

var TWITCH_OAUTH2_TOKEN string = TWITCH_OAUTH2_URL + "/token"
var TWITCH_OAUTH2_USERINFO string = TWITCH_OAUTH2_URL + "/userinfo"
