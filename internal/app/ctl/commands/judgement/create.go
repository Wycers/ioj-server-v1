package judgement

import (
	"github.com/infinity-oj/server/internal/app/ctl/service"
	"github.com/urfave/cli/v2"
)

type CreateJudgementCommand = cli.Command

func NewCreateJudgementsCommand(judgementService service.JudgementService) *ListJudgementCommand {
	return &ListJudgementCommand{
		Name:         "create",
		Aliases:      []string{"c"},
		Usage:        "create a judgement",
		UsageText:    "",
		Description:  "",
		ArgsUsage:    "",
		Category:     "",
		BashComplete: nil,
		Before:       nil,
		After:        nil,
		Action: func(c *cli.Context) error {
			submissionId := c.Uint64("submissionId")
			publicSpace := c.String("publicSpace")
			privateSpace := c.String("privateSpace")
			userSpace := c.String("userSpace")
			if err := judgementService.Create(submissionId, publicSpace, privateSpace, userSpace); err != nil {
				return err
			}
			return nil
		},
		OnUsageError: nil,
		Subcommands:  nil,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "publicSpace",
				Required: true,
				Aliases:  []string{"pub"},
				Usage:    "publicSpace of this judgement",
			},
			&cli.StringFlag{
				Name:     "privateSpace",
				Required: true,
				Aliases:  []string{"pri"},
				Usage:    "privateSpace of this judgement",
			},
			&cli.StringFlag{
				Name:     "userSpace",
				Required: true,
				Aliases:  []string{"usr"},
				Usage:    "userSpace of this judgement",
			}},
		SkipFlagParsing:        false,
		HideHelp:               false,
		Hidden:                 false,
		UseShortOptionHandling: false,
		HelpName:               "",
		CustomHelpTemplate:     "",
	}
}
