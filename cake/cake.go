package cake

import grpc "github.com/grpc/grpc-go"

type Server struct {
	*grpc.Server
}

func New() *Server {
	server := &Server{}

	return server
}
