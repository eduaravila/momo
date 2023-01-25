package utils

import "net/http"

func Handle(method string, route string, handler func(w http.ResponseWriter, r *http.Request), mux *http.ServeMux) {
	mux.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			handler(w, r)
			return
		}
		handler(w, r)
	})

}
