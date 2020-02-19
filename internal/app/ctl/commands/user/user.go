package user

import (
	"fmt"
	"github.com/google/wire"
	"github.com/urfave/cli/v2"
)

type Command struct {
	Command *cli.Command
}

var ProviderSet = wire.NewSet(NewUserCommand, NewCreateUserCommand)

func NewUserCommand(createUserCommand *CreateUserCommand) Command {
	var subCommands = []*cli.Command{
		createUserCommand.command,
	}
	return Command{Command: &cli.Command{
		Name:        "user",
		Aliases:     []string{"u"},
		Usage:       "options for user actions",
		UsageText:   "",
		Description: "",
		ArgsUsage:   "",
		Category:    "",
		BashComplete: func(c *cli.Context) {
			// This will complete if no args are passed
			if c.NArg() > 0 {
				return
			}
			for _, t := range subCommands {
				fmt.Println(t.Name)
			}
		},
		Before:                 nil,
		After:                  nil,
		Action:                 nil,
		OnUsageError:           nil,
		Subcommands:            subCommands,
		Flags:                  nil,
		SkipFlagParsing:        false,
		HideHelp:               false,
		Hidden:                 false,
		UseShortOptionHandling: false,
		HelpName:               "",
		CustomHelpTemplate:     "",
	}}
}
