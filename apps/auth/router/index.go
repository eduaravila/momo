package router

import (
	"net/http"

	"github.com/smollmegumin/momo/auth/router/twitch"
	"github.com/smollmegumin/momo/auth/util"
)

func Routes() *http.ServeMux {
	mux := util.NewHandler(http.NewServeMux())
	mux.Get("/oauth/twitch/callback", twitch.GetToken)
	return mux.GetServeMux()
}
