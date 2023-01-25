package routes

import (
	"net/http"

	"github.com/smollmegumin/level-TTS/listener/routes/twitch"
	"github.com/smollmegumin/level-TTS/listener/utils"
)

func Routes() *http.ServeMux {
	mux := utils.NewHandler(http.NewServeMux())
	mux.Get("/twitch/auth", twitch.GetToken)
	return mux.GetServeMux()
}
