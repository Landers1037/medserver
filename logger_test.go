package main

import "testing"

// test for logger
func TestLoggerInfo(t *testing.T) {
	logger := Logger{}
	_ = logger.InitLogger("", "", "", true)
	logger.Info("this is a test for info %s", "medserver")
	logger.Error("this is a test for error %s", "medserver")
}

func TestLoggerError(t *testing.T) {
	logger := Logger{}
	_ = logger.InitLogger("", "", "debug", true)
	logger.Info("this is a test for info %s", "medserver")
	logger.Error("this is a test for error %s", "medserver")
}

func TestLoggerNoTime(t *testing.T) {
	logger := Logger{}
	_ = logger.InitLogger("", "", "debug", false)
	logger.Info("this is a test for info %s", "medserver")
	logger.Error("this is a test for error %s", "medserver")
}
func TestLoggerFile(t *testing.T) {
	logger := Logger{}
	e := logger.InitLogger("test.log", "", "debug", false)
	t.Log(e)
	logger.Info("this is a test for info %s", "medserver")
	logger.Error("this is a test for error %s", "medserver")
}