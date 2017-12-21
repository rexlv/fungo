package base

import (
	"github.com/rexlv/fungo/cream/middleware"
	"github.com/rexlv/fungo/cream"
)

// Server base server
type Server struct {
	*cream.Server
}

// New returns new BaseServer
func New() *Server {
	server := &Server{
		Server: cream.New(),
	}

	server.Use(middleware.DefaultRecovery)

	return server
}
