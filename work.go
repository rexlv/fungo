package fungo

type Work interface {
	Work() error
	Stop()
	GracefulStop()
}
