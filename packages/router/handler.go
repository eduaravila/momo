package router

import "net/http"

// interface proxy that handles HTTP methods using http.ServeMux
type Handler interface {
	Get(route string, handler func(w http.ResponseWriter, r *http.Request))
	Post(route string, handler func(w http.ResponseWriter, r *http.Request))
	Put(route string, handler func(w http.ResponseWriter, r *http.Request))
	Patch(route string, handler func(w http.ResponseWriter, r *http.Request))
	Delete(route string, handler func(w http.ResponseWriter, r *http.Request))
	GetServeMux() *http.ServeMux
}

// Handler implementation
type handler struct {
	serveMux *http.ServeMux
}

// NewHandler returns a new instance of Handler
func NewHandler(mux *http.ServeMux) Handler {
	return &handler{serveMux: mux}
}

// Get handles GET method
func (h *handler) Get(route string, handler func(w http.ResponseWriter, r *http.Request)) {
	Handle(http.MethodGet, route, handler, h.serveMux)
}

// get the http.ServeMux
func (h *handler) GetServeMux() *http.ServeMux {
	return h.serveMux
}

// Post handles POST method
func (h *handler) Post(route string, handler func(w http.ResponseWriter, r *http.Request)) {
	Handle(http.MethodPost, route, handler, h.serveMux)
}

// Put handles PUT method
func (h *handler) Put(route string, handler func(w http.ResponseWriter, r *http.Request)) {
	Handle(http.MethodPut, route, handler, h.serveMux)
}

// Patch handles PATCH method
func (h *handler) Patch(route string, handler func(w http.ResponseWriter, r *http.Request)) {
	Handle(http.MethodPatch, route, handler, h.serveMux)
}

// Delete handles DELETE method
func (h *handler) Delete(route string, handler func(w http.ResponseWriter, r *http.Request)) {
	Handle(http.MethodDelete, route, handler, h.serveMux)
}

// Handle handles HTTP methods
func Handle(method string, route string, handler func(w http.ResponseWriter, r *http.Request), mux *http.ServeMux) {
	mux.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	})

}
