package main

import (
	"github.com/rexlv/fungo"
	"github.com/rexlv/fungo/cream"
)

type User struct {
	cream.RestHandler
}

func (u User) Query(c cream.Context) error {
	return nil
}

func (u User) Create(c cream.Context) error {

}

func (u User) Route() map[string]cream.HandlerFunc {
	return map[string]cream.HandlerFunc{
		"get /user/:id/info": u.Query,
		"post /user": u.Create,
	}
}

func main() {
	app := fungo.New()
	app.Init()
	app.Run()
}
