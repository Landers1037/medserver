package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func newServer(path string, mode bool) *gin.Engine {
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

	g.Static("/", path)
	return g
}

func runServer(path string, port int, test bool) error {
	server := newServer(path, test)

	return server.Run(fmt.Sprintf(":%d", port))
}