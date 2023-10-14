package routers

import (
	"scm/controllers"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func UserRouters(router *fasthttprouter.Router) {
	router.POST("/user/create", func(ctx *fasthttp.RequestCtx) {

	})
	router.POST("/user/createbulk", func(ctx *fasthttp.RequestCtx) {
		if controllers.CheckSession(ctx) != "" {
			controllers.AddUser(ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
}
