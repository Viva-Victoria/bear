package gorilla

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	r mux.Router
}

func (g *Router) Handle(path string, handler http.Handler, methods ...string) {
	g.r.Handle(path, handler).Methods(methods...)
}
