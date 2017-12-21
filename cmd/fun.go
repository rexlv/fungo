package main

import "strings"
import (
	"github.com/abiosoft/ishell"
	"github.com/docker/docker/cli/compose/types"
)

func main() {
	shell := ishell.New()

	shell.Println("Sample Interactive Shell")
	shell.SetPrompt(">>> ")

	// register a function for "greet" command.
    shell.AddCmd(&ishell.Cmd{
        Name: "greet",
        Help: "greet user",
        Func: func(c *ishell.Context) {
			c.Println("Hello", strings.Join(c.Args, " "))
			subshell := ishell.New()
			subshell.SetPrompt(">>> >>> ")
			subshell.AddCmd(&ishell.Cmd{
				Name: "hello",
				Help: "hello user",
				Func: func(subc *ishell.Context) {
					subc.Println("hello hello", strings.Join(subc.Args, " "))
				},
			})
			subshell.Run()
        },
    })

    // run shell
    shell.Run()
}
