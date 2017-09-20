package cream

// BaseServer
type BaseServer struct {
	*Server
}

func NewBaseServer() *BaseServer {
	return &BaseServer{
		Server: New(),
	}
}
