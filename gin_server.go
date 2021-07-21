package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func newServer(path string, mode, try bool) *gin.Engine {
	if mode {
		gin.SetMode(gin.DebugMode)
	}else {
		gin.SetMode(gin.ReleaseMode)
	}
	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %d %s \"%s\n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC1123),
			params.Method,
			params.Path,
			params.StatusCode,
			params.Latency,
			params.ErrorMessage)
	}))
	g.Use(MiddlewareTryFile(try))
	g.Static("/", path)

	return g
}

func runServer(path string, port int, test, try bool) error {
	server := newServer(path, test, try)

	return server.Run(fmt.Sprintf(":%d", port))
}