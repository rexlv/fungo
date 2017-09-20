package app

import (
	"fmt"

	"github.com/rexlv/fungo/cmd/command"
	"github.com/urfave/cli"
)

func init() {
	command.Register(
		cli.Command{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "create a fungo app",
			Action:  create,
			Subcommands: []cli.Command{
				{
					Name: "api",
				},
			},
		},
		cli.Command{},
	)
}

func create(c *cli.Context) error {
	fmt.Println("create")
	return nil
}
