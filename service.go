package fungo

import "net"

type Service interface {
	Serve(lis net.Listener) error
	Stop()
	GracefulStop()
}
