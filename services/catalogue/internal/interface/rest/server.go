package rest

import (
	"net/http"
)

type Server struct {
	srv     *http.Server
	handler http.Handler
}

func NewServer(addr string) *Server {
	srv := &http.Server{
		Addr: addr,
	}
	return &Server{
		srv: srv,
	}
}

func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}
