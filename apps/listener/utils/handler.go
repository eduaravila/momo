package utils

import "net/http"

func Handle(method string, route string, handler func(w http.ResponseWriter, r *http.Request), mux *http.ServeMux) {
	mux.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	})

}
