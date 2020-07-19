package submission

import (
	"fmt"

	"github.com/infinity-oj/server/internal/app/ctl/service"
	"github.com/urfave/cli/v2"
)

type DispatchSubmissionCommand = cli.Command

func NewDispatchSubmissionCommand(submissionService service.SubmissionService) *DispatchSubmissionCommand {
	return &DispatchSubmissionCommand{
		Name:         "dispatch",
		Aliases:      []string{"d"},
		Usage:        "dispatch judgement of a submission",
		UsageText:    "",
		Description:  "",
		ArgsUsage:    "",
		Category:     "",
		BashComplete: nil,
		Before:       nil,
		After:        nil,
		Action: func(c *cli.Context) error {
			//submissionId := c.String("submission ID")
			judgementId, err := submissionService.DispatchJudgement("88cfbbd7-678f-456c-a739-d1a2063ebf23")
			if err != nil {
				return err
			}
			fmt.Println(judgementId)
			return nil
		},
		OnUsageError: nil,
		Subcommands:  nil, Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "submission ID",
				Required: true,
				Aliases:  []string{"s", "sid"},
				Usage:    "submission ID",
			},
		},
		SkipFlagParsing:        false,
		HideHelp:               false,
		Hidden:                 false,
		UseShortOptionHandling: false,
		HelpName:               "",
		CustomHelpTemplate:     "",
	}
}
