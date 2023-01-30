package main

import (
	"net/http"

	"github.com/eduaravila/momo/apps/auth/router"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", router.Routes())
	http.ListenAndServe(":8080", mux)
}
