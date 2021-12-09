package router

import (
	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

type RouterGroup struct {
	*gin.RouterGroup
}

type HandleFunc func(*Context)

func New() *RouterGroup {
	return &RouterGroup{
		RouterGroup: &engine.RouterGroup,
	}
}

func Use(hs ...HandleFunc) {
	for _, h := range hs {
		engine.Use(handle(h))
	}
}

func UseGinHandleFunc(hs ...gin.HandlerFunc) {
	for _, h := range hs {
		engine.Use(h)
	}
}

func Run(addr string) {
	engine.Run(addr)
}

func (rg *RouterGroup) POST(relativePath string, h HandleFunc) {
	rg.RouterGroup.POST(relativePath, handle(h))
}

func (rg *RouterGroup) PUT(relativePath string, h HandleFunc) {
	rg.RouterGroup.PUT(relativePath, handle(h))
}

func (rg *RouterGroup) GET(relativePath string, h HandleFunc) {
	rg.RouterGroup.GET(relativePath, handle(h))
}

func (rg *RouterGroup) GinGET(relativePath string, h ...gin.HandlerFunc) {
	rg.RouterGroup.GET(relativePath, h...)
}

func (rg *RouterGroup) DELETE(relativePath string, h HandleFunc) {
	rg.RouterGroup.DELETE(relativePath, handle(h))
}

func (rg *RouterGroup) Group(relativePath string) *RouterGroup {
	return &RouterGroup{
		RouterGroup: rg.RouterGroup.Group(relativePath),
	}
}

func (rg *RouterGroup) Handle(method string, relativePath string, h gin.HandlerFunc) {
	rg.RouterGroup.Handle(method, relativePath, h)
}
func handle(h HandleFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := &Context{
			Context: ctx,
		}
		h(c)
	}
}

func (rg *RouterGroup) Use(handlers ...HandleFunc) {
	for _, h := range handlers {
		rg.RouterGroup.Use(handle(h))
	}
}

func (rg *RouterGroup) GinUse(handlers ...gin.HandlerFunc) {
	rg.RouterGroup.Use(handlers...)
}

func init() {
	engine = gin.New()
	Use(Log)
}
