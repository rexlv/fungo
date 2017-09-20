package fungo

import (
	"net"
	"sync"
)

// App struct
type App struct {
	srvs      map[string]Service
	jobs      map[string]Work
	listeners map[string]net.Listener
	deferFns  []func()
}

// New new app
func New() *App {
	return &App{
		srvs:     make(map[string]Service),
		jobs:     make(map[string]Work),
		deferFns: make([]func(), 0),
	}
}

// Init init application
func (app *App) Init() {
}

func (app *App) Defer(fn func()) {
	app.deferFns = append([]func(){fn}, app.deferFns...)
}

// Run app
func (app *App) Run() {
	var wg sync.WaitGroup
	for name, srv := range app.srvs {
		wg.Add(1)
		srv.Mux()
		lis := app.listeners[name]
		go func(lis net.Listener, srv Service) {
			defer wg.Done()
			srv.Serve(lis)
		}(lis, srv)
	}

	wg.Wait()
	for _, fn := range app.deferFns {
		fn()
	}
}

func (app *App) run() {

}
