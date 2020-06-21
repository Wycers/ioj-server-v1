package file

import (
	"fmt"

	"github.com/infinity-oj/server/internal/app/ctl/service"

	"github.com/google/wire"
	"github.com/urfave/cli/v2"
)

type Command struct {
	Command *cli.Command
}

var ProviderSet = wire.NewSet(NewFileCommand, NewCreateDirectoryCommand)

func NewFileCommand(createDirectoryCommand *CreateDirectoryCommand, fileService service.FileService) Command {
	var subCommands = []*cli.Command{
		createDirectoryCommand.command,
		NewUploadCommand(fileService),
	}
	return Command{Command: &cli.Command{
		Name:        "file",
		Aliases:     []string{"f"},
		Usage:       "options for file actions",
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
