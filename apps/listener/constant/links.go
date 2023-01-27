package constant

import (
	"os"
)

// TWITCH_CLIENT_ID is the client id for the twitch app

var TWITCH_APPLICATION_CLIEND_ID string = os.Getenv("TWITCH_APPLICATION_CLIEND_ID")

// TWITCH_CLIENT_SECRET is the client secret for the twitch app

var DASHBOARD_APP_URL string = os.Getenv("DASHBOARD_APP_URL")

// TWITCH_TOKEN_URL app client id

var TWITCH_OAUTH2_URL string = os.Getenv("TWITCH_OAUTH2_URL")

// TWITCH_REDIRECT_URI is the url to redirect to after auth

var TWITCH_CLIENT_SECRET string = os.Getenv("TWITCH_CLIENT_SECRET")
