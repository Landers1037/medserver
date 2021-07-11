package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"time"
)

// actions group of cmds
func initAction(c *cli.Context) error {
	// get flags
	path := c.String("path")
	port := c.Int("port")
	logFile := c.String("log")
	testMode := c.Bool("test")
	daemon := c.Bool("daemon")
	withtime := c.Bool("withtime")

	var logg Logger
	var level string
	if testMode {
		level = "debug"
	}else {
		level = "info"
	}
	_ = logg.InitLogger(logFile, Name, level, withtime)
	showInfo(path, logFile, port, testMode, daemon)

	err := runServer(path, port, testMode)
	if err != nil {
		logg.Error("med exit with: %s", err.Error())
	}
	return err
}

func showInfo(path, log string, port int, test, daemon bool) {
	fmt.Printf("~%s running: %d\n", Name, port)
	fmt.Printf("version: %s\n" +
		"usage: %s\n" +
		"serve time: %s\n" +
		"static path: %s\n" +
		"logfile: %s\n" +
		"isTest: %v\n" +
		"isDaemon: %v\n", Version, Usage, time.Now().UTC(), path, log, test, daemon)
}