package fungo

import (
	"net"
	"reflect"
	"fmt"
	"github.com/rexlv/fungo/cream"
)

// Service interface
type Service interface {
	GracefulStop()
	Serve(lis net.Listener) error
	Stop()

	Label() string
	Mux()
}

// Serve serve
func (app *App) Serve(srv Service, opts ...ServiceOption) {
	var options ServiceOptions

	for _, opt := range opts {
		opt(&options)
	}

	typ := reflect.TypeOf(srv)
	val := reflect.ValueOf(srv).Elem()
	if typ.Kind() != reflect.Ptr {
		panic("pointer required")
	}

	if typ.Elem().Kind() != reflect.Struct {
		panic("struct pointer required")
	}

	switch typ.Elem().Field(0).Type {
	case reflect.TypeOf(new(cream.Server)):
		fmt.Println("cream server")
		val.Field(0).Set(reflect.ValueOf(cream.New()))
	default:
		panic("cream server at field required")
	}

	fmt.Println("type", typ, val, srv)


	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}

	app.srvs[srv.Label()] = srv
	app.listeners[srv.Label()] = listener
}

