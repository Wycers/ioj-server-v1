package file

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/Infinity-OJ/Server/internal/app/ctl/service"
	"github.com/urfave/cli/v2"
)

func uploadFile(fileService service.FileService, localFilePath, space, remoteDir string) (err error) {
	fmt.Println(localFilePath)
	fmt.Println(space)
	fmt.Println(remoteDir)

	dat, err := ioutil.ReadFile(localFilePath)
	if err != nil {
		return err
	}

	err = fileService.CreateFile(space, path.Join(remoteDir, path.Base(localFilePath)), dat)
	if err != nil {
		return
	}
	return
}

func uploadDirectory(fileService service.FileService, base, localDir, space, remoteDir string) (err error) {
	files, err := ioutil.ReadDir(path.Join(base, localDir))
	if err != nil {
		return
	}

	err = fileService.CreateDirectory(space, localDir)
	if err != nil {
		return
	}

	for _, f := range files {
		if f.IsDir() {
			if err = uploadDirectory(fileService, base, path.Join(localDir, f.Name()), space, path.Join(remoteDir, f.Name())); err != nil {
				return
			}
		} else {
			if err = uploadFile(fileService, path.Join(base, localDir, f.Name()), space, remoteDir); err != nil {
				return
			}
		}
	}
	return
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
			s := c.String("space")
			p := c.String("path")
			r := c.Bool("recursive")
			if r {
				if err := uploadDirectory(fileService, p, "", s, ""); err != nil {
					return err
				}
			} else {
				if err := uploadFile(fileService, p, s, ""); err != nil {
					return err
				}
			}

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
