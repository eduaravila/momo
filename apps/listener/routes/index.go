package routes

import (
	"net/http"

	"github.com/smollmegumin/level-TTS/listener/routes/twitch"
	"github.com/smollmegumin/level-TTS/listener/utils"
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	utils.Handle("GET", "/twitch/auth", twitch.GetToken, mux)
	return mux
}
