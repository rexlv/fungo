package fungo

import (
	"fmt"
	"net"
	"reflect"
)

// Service interface
type Service interface {
	Serve(lis net.Listener) error
	Stop()
	GracefulStop()

	Mux()
}

func clone(i interface{}) interface{} {
	// Wrap argument to reflect.Value, dereference it and return back as interface{}
	return reflect.Indirect(reflect.ValueOf(i)).Interface()
}

// Serve serve
func (app *App) Serve(srv Service, addr string) {
	t := reflect.ValueOf(srv).Elem().Type()
	//v := reflect.New(reflect.TypeOf(srv).Elem())

	fmt.Println("t: ", t)

	app.srvs[addr] = srv
}
