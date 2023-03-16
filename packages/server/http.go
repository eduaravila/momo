package server

import (
	"log"
	"net/http"
)

func RunHTTPServer(
	prefix,
	addrs string,
	createHandler func() http.Handler) {
	rootRouter := http.NewServeMux()
	rootRouter.Handle(prefix, createHandler())

	if err := http.ListenAndServe(addrs, rootRouter); err != nil {
		log.Fatal(err)
	}

}
