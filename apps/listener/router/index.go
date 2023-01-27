package router

import (
	"net/http"

	"github.com/smollmegumin/momo/listener/router/twitch"
	"github.com/smollmegumin/momo/listener/util"
)

func Routes() *http.ServeMux {
	mux := util.NewHandler(http.NewServeMux())
	mux.Get("/oauth/twitch/callback", twitch.GetToken)
	return mux.GetServeMux()
}
