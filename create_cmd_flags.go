package main

import "github.com/urfave/cli/v2"

// flags of cmd

// flags
func getFlagsGroup() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "use",
			Aliases: []string{"u"},
			Usage:   "set default engine [gin|http]",
			Value:   "gin",
		},
		&cli.StringFlag{
			Name:    "path",
			Aliases: []string{"p", "static"},
			Usage:   "set default static file path",
			Value:   ".",
		},
		&cli.IntFlag{
			Name:    "port",
			Aliases: []string{"po"},
			Usage:   "set default listen port",
			Value:   5000,
		},
		&cli.StringFlag{
			Name:    "log",
			Aliases: []string{"l", "logfile"},
			Usage:   "set default logfile",
			Value:   "",
		},
		&cli.BoolFlag{
			Name:    "withtime",
			Aliases: []string{"enable", "time"},
			Usage:   "log with current time",
			Value:   false,
		},
		&cli.BoolFlag{
			Name:    "test",
			Aliases: []string{"t", "testit"},
			Usage:   "run in a test mode",
			Value:   false,
		},
		&cli.BoolFlag{
			Name:    "daemon",
			Aliases: []string{"d"},
			Usage:   "run in a daemon mode",
			Value:   false,
		},
	}
}

// cmds
func getCommandsGroup() []*cli.Command {
	return []*cli.Command{}
}