package url

import (
	"os"
)

// DASHBOARD_APP_URL app client id
var DASHBOARD_APP_URL string = os.Getenv("DASHBOARD_APP_URL")

// TWITCH_TOKEN_URL app client id
var TWITCH_API_URL string = os.Getenv("TWITCH_OAUTH2_URL")
