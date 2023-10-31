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
}
