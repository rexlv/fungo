package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	_ "github.com/rexlv/fungo/cmd/command/app"
)

func main() {
	app := cli.NewApp()
	app.Name = "fun"
	app.Usage = "make an explosive entrance"
	app.Action = func(c *cli.Context) error {
		fmt.Println("fun! I say!")
		return nil
	}

	app.Run(os.Args)
}
