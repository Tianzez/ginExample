package router

import (
	"time"

	"github.com/cihub/seelog"
)

func Log(c *Context) {
	start := time.Now()
	seelog.Tracef("%#v", c.Request)
	c.Context.Next()
	latency := time.Since(start)
	seelog.Infof("%15s %6s %3d %13v %s", c.ClientIP(), c.Request.Method, c.Writer.Status(), latency, c.Request.URL.Path)
}

func Recover(c *Context) {
	defer func() {
		if r := recover(); r != nil {
			c.Failure(r)
			c.Abort()
		}
	}()
	c.Context.Next()
}
