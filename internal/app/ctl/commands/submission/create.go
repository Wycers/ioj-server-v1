package submission

import (
	"github.com/infinity-oj/server/internal/app/ctl/service"
	"github.com/urfave/cli/v2"
)

type CreateSubmissionCommand = cli.Command

func NewCreateSubmissionCommand(submissionService service.SubmissionService) *CreateSubmissionCommand {
	return &CreateSubmissionCommand{
		Name:         "create",
		Aliases:      []string{"c"},
		Usage:        "create a new submission",
		UsageText:    "",
		Description:  "",
		ArgsUsage:    "",
		Category:     "",
		BashComplete: nil,
		Before:       nil,
		After:        nil,
		Action: func(c *cli.Context) error {
			submitterId := c.Uint64("submitterId")
			problemId := c.String("problemId")
			userSpace := c.String("fileSpace")
			if err := submissionService.Create(submitterId, problemId, userSpace); err != nil {
				return err
			}
			return nil
		},
		OnUsageError: nil,
		Subcommands:  nil, Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "submitterId",
				Required: true,
				Aliases:  []string{"u", "user"},
				Usage:    "submitter ID of the submission",
			},
			&cli.StringFlag{
				Name:     "problemId",
				Required: true,
				Aliases:  []string{"p"},
				Usage:    "problem ID of the submission",
			},
			&cli.StringFlag{
				Name:     "fileSpace",
				Required: true,
				Aliases:  []string{"f"},
				Usage:    "file space containing submitting code",
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
