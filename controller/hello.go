package controller

import "github.com/Tianzez/ginExample/lib/router"

func SayHello(c *router.Context) {
	c.Success("hello")
}
