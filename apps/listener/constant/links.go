package constant

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// TWITCH_CLIENT_ID is the client id for the twitch app
var TWITCH_OAUTH2_URL string

// TWITCH_CLIENT_SECRET is the client secret for the twitch app
var TWITCH_CLIENT_SECRET string

// TWITCH_TOKEN_URL app client id
var TWITCH_APPLICATION_CLIEND_ID string

// TWITCH_REDIRECT_URI is the url to redirect to after auth
var DASHBOARD_APP_URL string

func init() {
	env := os.Getenv("ENV")
	fmt.Println("env: " + env)
	if env == "" {
		env = "development"
	}

	err := godotenv.Load(".env." + env + ".local")
	if err != nil {
		log.Fatal("Error loading .env file" + err.Error())

	}
	TWITCH_APPLICATION_CLIEND_ID = os.Getenv("TWITCH_APPLICATION_CLIEND_ID")
	DASHBOARD_APP_URL = os.Getenv("DASHBOARD_APP_URL")
	TWITCH_OAUTH2_URL = os.Getenv("TWITCH_OAUTH2_URL")
	TWITCH_CLIENT_SECRET = os.Getenv("TWITCH_CLIENT_SECRET")

}
