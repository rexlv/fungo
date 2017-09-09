package fungo

import (
	"sync"
)

type App struct {
	srvs map[string]Service
	jobs map[string]Job
}

func New() *App {
	return &App{
		srvs: make(map[string]Service),
		jobs: make(map[string]Job),
	}
}

// Init init application
func (app *App) Init() {
}

func (app *App) Run() {
	var wg sync.WaitGroup
	for name, srv := range app.srvs {
		wg.Add(1)
		go func(srv Service) {
			srv.Run()
		}(srv)
	}

	wg.Wait()
}

func (app *App) run() {

}
