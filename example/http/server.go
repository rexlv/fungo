package http

import "github.com/rexlv/fungo/cream"

type Server struct {
	cream.Server
}

func (s *Server) Mux() {
	s.GET("/ping", pong)
}

func pong(c cream.Context) error {
	return nil
}
