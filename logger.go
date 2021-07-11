package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

// Logger is a simple format logger
const (
	INFO = "[INFO] "
	ERROR = "[ERROR] "
	PREFIX = "medserver - "
	DEFAULTLEVEL = "info"
)

type Logger struct {
	LogFile string
	LogPrefix string
	EnableTimeStamp bool
	LogLevel string // info debug
	LoggerIns *log.Logger
}

func (logger *Logger)InitLogger(file, prefix, level string, enable bool) error {
	// check log file
	var writer io.Writer
	if _, e := os.Stat(file); os.IsNotExist(e) && file != "" {
		writer = os.Stdout
	}
	f, e := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if e != nil {
		return errors.New("log file can't create")
	}else {
		writer = io.MultiWriter(f, os.Stdout)
	}

	if prefix == "" {
		prefix = PREFIX
	}

	if enable {
		logger.LoggerIns = log.New(writer, prefix, log.Ldate|log.Ltime|log.Lshortfile)
	}else {
		logger.LoggerIns = log.New(writer, prefix, log.Lshortfile)
	}
	logger.LogFile = file
	logger.LogPrefix = prefix
	logger.EnableTimeStamp = enable
	logger.LogLevel = level
	return nil
}

// info
func (logger *Logger)Info(fmtStr string, args ...interface{}) {
	if len(args) == 0 {
		logger.LoggerIns.Println(INFO + fmtStr)
	}else {
		logger.LoggerIns.Println(INFO + fmt.Sprintf(fmtStr, args...))
	}
}

// error
func (logger *Logger)Error(fmtStr string, args ...interface{}) {
	if logger.LogLevel == "debug" {
		if len(args) == 0 {
			logger.LoggerIns.Println(ERROR + fmtStr)
		}else {
			logger.LoggerIns.Println(ERROR + fmt.Sprintf(fmtStr, args...))
		}
	}
}