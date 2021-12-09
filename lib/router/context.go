package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"runtime/debug"

	"github.com/Tianzez/ginExample/lib/errors"
	"github.com/Tianzez/ginExample/utils"
	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Context struct {
	*gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg"`
}

func (r *Response) String() string {
	ss, _ := json.Marshal(r)
	return string(ss)
}

func (c *Context) SuccessWithMsg(code int, msg string, data interface{}) {
	c.EchoResponse(code, msg, data)
}

func (c *Context) Success(data interface{}) {
	c.EchoResponse(200, "", data)
}

func (c *Context) Failure(r interface{}) {
	if e, ok := r.(*errors.Error); ok {
		c.EchoResponse(e.Code, e.Msg, nil)
		return
	}
	if _, ok := r.(runtime.Error); ok {
		log.Debugf("Internal error, stack:%s\n", debug.Stack())
	}
	if e, ok := r.(error); ok {
		if e.Error() == gorm.ErrRecordNotFound.Error() {
			c.EchoResponse(400, "抱歉，您访问的资源不存在", nil)
			return
		}
		c.EchoResponse(400, e.Error(), nil)
		return
	}

	if s, ok := r.(string); ok {
		c.EchoResponse(400, s, nil)
		return
	}
	c.EchoResponse(500, "Internal error", nil)
}

func (c *Context) EchoResponse(code int, msg string, data interface{}) {
	r := Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	c.Echo(code, r.String())
}

func (c *Context) Echo(code int, s string) {
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	httpCode := http.StatusOK
	if code != 200 {
		httpCode = code
	}
	c.String(httpCode, s)
}

func (c *Context) Param(key string) utils.String {
	v := c.Context.Param(key)
	return utils.String(v)
}

func (c *Context) GetID() utils.String {
	return c.Param("id")
}

func (c *Context) Query(key string) utils.String {
	v := c.Context.Query(key)
	return utils.String(v)
}

func (c *Context) ReadStringNonNil(key string) string {
	v := c.Context.Query(key)
	if v == "" {
		errors.ThrowIfNotNil(fmt.Errorf("%s is a required parameter", key), http.StatusBadRequest)
	}
	return utils.String(v).String()
}

func (c *Context) ReadIntWithDef(key string, def int) int {
	v := c.DefaultQuery(key, fmt.Sprintf("%d", def))
	return utils.String(v).Int()
}
