package main

import (
	. "github.com/rexlv/fungo"
	"github.com/rexlv/fungo/example/http"
	"fmt"
)


func main() {
	app := New()
	app.Serve(
		new(http.Server),
		Addr(":9090"),
	)

	app.Defer(func(){
		fmt.Println("defer")
	})

	app.Defer(func(){
		fmt.Println("defer2")
	})

	app.Run()
}
