package router

import (
	"github.com/Tianzez/ginExample/controller"
	"github.com/Tianzez/ginExample/lib/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
)

func Init() {
	router.UseGinHandleFunc(gzip.Gzip(gzip.DefaultCompression))
	router.UseGinHandleFunc(cors.Default())
	router.Use(router.Recover)

	r := router.New().Group("/api/v1")
	r.GET("/hello", controller.SayHello)

	router.Run(":8901")

}
