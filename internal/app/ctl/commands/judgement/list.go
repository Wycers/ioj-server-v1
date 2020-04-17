package judgement

import (
	"github.com/Infinity-OJ/Server/internal/app/ctl/service"
	"github.com/urfave/cli/v2"
)

type ListJudgementCommand = cli.Command

func NewListJudgementsCommand(judgementService service.JudgementService) *ListJudgementCommand {
	return &ListJudgementCommand{
		Name:         "list",
		Aliases:      []string{"c"},
		Usage:        "list all judgements",
		UsageText:    "",
		Description:  "",
		ArgsUsage:    "",
		Category:     "",
		BashComplete: nil,
		Before:       nil,
		After:        nil,
		Action: func(c *cli.Context) error {
			if err := judgementService.List(); err != nil {
				return err
			}
			return nil
		},
		OnUsageError:           nil,
		Subcommands:            nil,
		Flags:                  []cli.Flag{},
		SkipFlagParsing:        false,
		HideHelp:               false,
		Hidden:                 false,
		UseShortOptionHandling: false,
		HelpName:               "",
		CustomHelpTemplate:     "",
	}
}
