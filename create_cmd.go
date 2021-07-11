package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

const (
	Name = "med"
	Usage = "simple static server"
	UsageText = "med [options] [cmds] args"
	Des = "med is a simple static-file server"
	CopyRight = "renj.io"
)

var Version = fmt.Sprintf("build%d%d%d", time.Now().Year(),time.Now().Month(), time.Now().Day())
var Author = []*cli.Author{
	{
		Name: "Landers",
		Email: "mail@renj.io",
	},
}

// InitCMD create cmd app by cli
func InitCMD() error {
	app := &cli.App{
		Name:                   Name,
		HelpName:               Name,
		Usage:                  Usage,
		UsageText:              UsageText,
		ArgsUsage:              "",
		Version:                Version,
		Description:            Des,
		Commands:               getCommandsGroup(),
		Flags:                  getFlagsGroup(),
		EnableBashCompletion:   false,
		HideHelp:               false,
		HideHelpCommand:        false,
		HideVersion:            false,
		BashComplete:           nil,
		Before:                 nil,
		After:                  nil,
		Action: 				initAction,
		CommandNotFound:        cmdNotFound,
		OnUsageError:           usageError,
		Compiled:               time.Time{},
		Authors:                Author,
		Copyright:              CopyRight,
		Reader:                 nil,
		Writer:                 nil,
		ErrWriter:              nil,
		ExitErrHandler:         nil,
		Metadata:               nil,
		ExtraInfo:              nil,
		CustomAppHelpTemplate:  "",
		UseShortOptionHandling: false,
	}
	// init cmd
	return app.Run(os.Args)
}

func cmdNotFound(c *cli.Context, str string) {
	fmt.Printf("[medserver] cmd not found: %s\n", str)
}

func usageError(c *cli.Context,  err error, isSubcommand bool) error {
	fmt.Printf("[medserver] usage error for %s\n", err.Error())
	return err
}