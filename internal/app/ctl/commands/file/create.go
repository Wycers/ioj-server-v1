package file

import (
	"github.com/Infinity-OJ/Server/internal/app/ctl/service"
	"github.com/urfave/cli/v2"
)

type CreateDirectoryCommand struct {
	command *cli.Command
}

func NewCreateDirectoryCommand(fileService service.FileService) *CreateDirectoryCommand {
	return &CreateDirectoryCommand{command: &cli.Command{
		Name:         "create",
		Aliases:      []string{"c"},
		Usage:        "create a new directory",
		UsageText:    "",
		Description:  "",
		ArgsUsage:    "",
		Category:     "",
		BashComplete: nil,
		Before:       nil,
		After:        nil,
		Action: func(c *cli.Context) error {
			//fmt.Println("new task template: ", c.Args().First())
			//name := c.String("name")
			fileService.Create("qwq")
			return nil
		},
		OnUsageError: nil,
		Subcommands:  nil, Flags: nil,
		SkipFlagParsing:        false,
		HideHelp:               false,
		Hidden:                 false,
		UseShortOptionHandling: false,
		HelpName:               "",
		CustomHelpTemplate:     "",
	}}
}
