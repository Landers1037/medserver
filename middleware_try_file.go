/*
Name: medserver
File: middleware_try_file.go
Author: Landers
*/

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// MiddlewareTryFile 对于webpack机制的静态路由增加try机制
func MiddlewareTryFile(try bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 根据用户传入的static path解析内部的html
		// 转发所有的404到index页面
		c.Next()
		if try {
			if c.Writer.Status() == http.StatusNotFound {
				c.File("index.html")
			}
		}
	}
}
