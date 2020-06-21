package submission

import (
	"fmt"

	"github.com/infinity-oj/server/internal/app/ctl/service"

	"github.com/google/wire"
	"github.com/urfave/cli/v2"
)

type Command struct {
	Command *cli.Command
}

var ProviderSet = wire.NewSet(NewSubmissionCommand)

func NewSubmissionCommand(submissionService service.SubmissionService) Command {
	var subCommands = []*cli.Command{
		NewCreateSubmissionCommand(submissionService),
	}
	return Command{Command: &cli.Command{
		Name:        "submission",
		Aliases:     []string{"s"},
		Usage:       "options for submission actions",
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
