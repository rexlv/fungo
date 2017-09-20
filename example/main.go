package main

import (
	"github.com/rexlv/fungo"
	"github.com/rexlv/fungo/example/http"
)

func main() {
	app := fungo.New()
	app.Init()

	app.Serve(&http.Server{}, ":9999")
	app.Run()
}
