package user

import (
	"github.com/infinity-oj/server/internal/app/ctl/service"
	"github.com/urfave/cli/v2"
)

type CreateUserCommand struct {
	command *cli.Command
}

func NewCreateUserCommand(userService service.UserService) *CreateUserCommand {
	return &CreateUserCommand{command: &cli.Command{
		Name:         "create",
		Aliases:      []string{"c"},
		Usage:        "create a new user",
		UsageText:    "",
		Description:  "",
		ArgsUsage:    "",
		Category:     "",
		BashComplete: nil,
		Before:       nil,
		After:        nil,
		Action: func(c *cli.Context) error {
			//fmt.Println("new task template: ", c.Args().First())
			username := c.String("username")
			password := c.String("password")
			email := c.String("email")
			userService.Create(username, password, email)
			return nil
		},
		OnUsageError: nil,
		Subcommands:  nil, Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "username",
				Required: true,
				Aliases:  []string{"u", "user"},
				Usage:    "username for new user",
			},
			&cli.StringFlag{
				Name:     "email",
				Required: true,
				Aliases:  []string{"e"},
				Usage:    "email for new user",
			},
			&cli.StringFlag{
				Name:     "password",
				Required: true,
				Aliases:  []string{"p"},
				Usage:    "password for new user",
			},
		},
		SkipFlagParsing:        false,
		HideHelp:               false,
		Hidden:                 false,
		UseShortOptionHandling: false,
		HelpName:               "",
		CustomHelpTemplate:     "",
	}}
}
