package bear

import (
	"errors"
	"net/http"
)

var (
	ErrTooManyErrors = errors.New("too many errors occurred")
)

type Handler func(ctx Context, r *http.Request, w http.ResponseWriter)

type Server struct {
	server    http.Server
	router    Router
	logger    Logger
	maxErrors int
}

func (s *Server) Start() {
	go s.Run()
}

func (s *Server) Handle(path string, handler Handler, methods ...string) {
	s.router.Handle(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler(ReadContext(r), r, w)
	}), methods...)
}

func (s *Server) Run() error {
	count := 0
	for {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Warn("ListenAndServe failed", err)

			count++
			if count > s.maxErrors {
				return ErrTooManyErrors
			}
		}
	}
}

func (s *Server) Stop() error {
	return s.server.Close()
}
