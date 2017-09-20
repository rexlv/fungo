package command

import "github.com/urfave/cli"

var commands = map[string]cli.Command{}

// Register register command
func Register(cmds ...cli.Command) {
	for _, cmd := range cmds {
		commands[cmd.FullName()] = cmd
	}
}
