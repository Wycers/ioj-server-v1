package ctl

import (
	"github.com/Infinity-OJ/Server/internal/app/ctl/commands/user"
	"github.com/google/wire"
	"github.com/urfave/cli/v2"
	"time"
)

func NewApp(userCommand user.Command) *cli.App {
	app := &cli.App{
		Name:                   "",
		HelpName:               "",
		Usage:                  "",
		UsageText:              "",
		ArgsUsage:              "",
		Version:                "",
		Description:            "",
		Commands:               []*cli.Command{userCommand.Command},
		Flags:                  nil,
		EnableBashCompletion:   false,
		HideHelp:               false,
		HideVersion:            false,
		BashComplete:           nil,
		Before:                 nil,
		After:                  nil,
		Action:                 nil,
		CommandNotFound:        nil,
		OnUsageError:           nil,
		Compiled:               time.Time{},
		Authors:                nil,
		Copyright:              "",
		Writer:                 nil,
		ErrWriter:              nil,
		ExitErrHandler:         nil,
		Metadata:               nil,
		ExtraInfo:              nil,
		CustomAppHelpTemplate:  "",
		UseShortOptionHandling: false,
	}

	return app
}

var ProviderSet = wire.NewSet(NewApp)
