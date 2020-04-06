package file

import (
	"fmt"
	"io/ioutil"

	"github.com/Infinity-OJ/Server/internal/app/ctl/service"
	"github.com/urfave/cli/v2"
)

func uploadFile(fileService service.FileService, localFilePath, space, remoteFilePath string) (err error) {
	fmt.Println(localFilePath)
	fmt.Println(space)
	fmt.Println(remoteFilePath)

	dat, err := ioutil.ReadFile(localFilePath)
	if err != nil {
		return err
	}

	err = fileService.CreateFile(space, remoteFilePath+"XD.txt", dat)
	if err != nil {
		return
	}
	return
}

func uploadDirectory(fileService *service.FileService, localDir, remoteDir string) {

}

func NewUploadCommand(fileService service.FileService) *cli.Command {
	return &cli.Command{
		Name:         "upload",
		Aliases:      []string{"up"},
		Usage:        "upload a file or directory",
		UsageText:    "",
		Description:  "",
		ArgsUsage:    "",
		Category:     "",
		BashComplete: nil,
		Before:       nil,
		After:        nil,

		Action: func(c *cli.Context) error {
			//fmt.Println("new task template: ", c.Args().First())
			fmt.Println("?")
			s := c.String("space")
			p := c.String("path")
			r := c.Bool("recursive")
			if r {
				//uploadDirectory(p)
			} else {
				if err := uploadFile(fileService, p, s, ""); err != nil {
					//fmt.Println(err.Error())
					fmt.Println("???")
					return err
				}
			}
			fmt.Println(r)

			return nil
		},

		OnUsageError: nil,
		Subcommands:  nil, Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "space",
				Required: true,
				Aliases:  []string{"s"},
				Usage:    "target file space you want to upload",
			},
			&cli.StringFlag{
				Name:     "path",
				Required: true,
				Aliases:  []string{"p"},
				Usage:    "file or directory you want to upload",
			},
			&cli.BoolFlag{
				Name:     "recursive",
				Required: false,
				Aliases:  []string{"r", "R"},
				Usage:    "upload directories and their contents recursively",
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
