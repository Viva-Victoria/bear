package bear

import (
	"net/http"
)

type Router interface {
	Handle(path string, handler http.Handler, methods ...string)
}
