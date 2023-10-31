package routers

import (
	"scm/controllers"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func Access_MenuRouters(router *fasthttprouter.Router) {
	router.POST("/admin/menu/add", func(ctx *fasthttp.RequestCtx) {
		if controllers.CheckSession(ctx) != "" {
			controllers.AddMenu(ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("/admin/access1/add", func(ctx *fasthttp.RequestCtx) {
		if controllers.CheckSession(ctx) != "" {
			controllers.AddAccess1(ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("/admin/access2/add", func(ctx *fasthttp.RequestCtx) {
		if controllers.CheckSession(ctx) != "" {
			controllers.AddAccess1(ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
}
