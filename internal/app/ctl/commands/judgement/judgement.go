package judgement

import (
	"fmt"

	"github.com/Infinity-OJ/Server/internal/app/ctl/service"

	"github.com/google/wire"
	"github.com/urfave/cli/v2"
)

type Command struct {
	Command *cli.Command
}

var ProviderSet = wire.NewSet(NewJudgementCommand)

func NewJudgementCommand(judgementService service.JudgementService) Command {
	var subCommands = []*cli.Command{
		NewListJudgementsCommand(judgementService),
	}
	return Command{Command: &cli.Command{
		Name:        "judgement",
		Aliases:     []string{"s"},
		Usage:       "options for judgement actions",
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
