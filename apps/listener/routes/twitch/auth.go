package twitch

import (
	"net/http"
)

type TwithRouter struct {
}

func NewTwitchRouter() *TwithRouter {
	return &TwithRouter{}
}

func (t *TwithRouter) GET(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
