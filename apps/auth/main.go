package main

import (
	"net/http"

	"github.com/smollmegumin/momo/auth/router"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", router.Routes())
	http.ListenAndServe(":8080", mux)

}
