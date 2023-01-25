package twitch

import (
	"net/http"
)

func GetToken(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
