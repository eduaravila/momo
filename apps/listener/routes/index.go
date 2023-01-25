package routes

import (
	"fmt"
	"net/http"
)

type FallbackMethods struct {
	methods any
}

func (f FallbackMethods) GET(w http.ResponseWriter, r *http.Request) {
	if c, ok := f.methods.(GET); ok {
		c.GET(w, r)
		return
	}
	defer r.Body.Close()
	w.WriteHeader(http.StatusBadGateway)
	w.Write([]byte("This is a fallback"))
}
func (f FallbackMethods) POST(w http.ResponseWriter, r *http.Request) {
	if c, ok := f.methods.(POST); ok {
		c.POST(w, r)
	}

}
func (f FallbackMethods) PUT(w http.ResponseWriter, r *http.Request) {
	if c, ok := f.methods.(PUT); ok {
		c.PUT(w, r)
	}
}
func (f FallbackMethods) DELETE(w http.ResponseWriter, r *http.Request) {
	if c, ok := f.methods.(DELETE); ok {
		c.DELETE(w, r)
	}
}

type GET interface {
	GET(w http.ResponseWriter, r *http.Request)
}

type POST interface {
	POST(w http.ResponseWriter, r *http.Request)
}

type PUT interface {
	PUT(w http.ResponseWriter, r *http.Request)
}

type DELETE interface {
	DELETE(w http.ResponseWriter, r *http.Request)
}

type Methods interface {
	GET(w http.ResponseWriter, r *http.Request)
	POST(w http.ResponseWriter, r *http.Request)
	PUT(w http.ResponseWriter, r *http.Request)
	DELETE(w http.ResponseWriter, r *http.Request)
}

type Router interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type router struct {
	Methods
}

func (ru *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ru.GET(w, r)
	case "POST":
		ru.POST(w, r)
	case "PUT":
		ru.PUT(w, r)
	case "DELETE":
		ru.DELETE(w, r)
	}
}

func NewRoute(m any) Router {
	fmt.Println(m)
	return &router{
		Methods: &FallbackMethods{m},
	}
}
