package routers

import (
	"scm/controllers"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func Access_MenuRouters(router *fasthttprouter.Router) {
	router.POST("/api/v1/admin/menu/add", func(ctx *fasthttp.RequestCtx) {
		_, _, errMsg := controllers.CheckSession(ctx)
		if errMsg == "" {
			controllers.AddMenu(ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("/api/v1/admin/access1/add", func(ctx *fasthttp.RequestCtx) {
		_, _, errMsg := controllers.CheckSession(ctx)
		if errMsg == "" {
			controllers.AddAccess(ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
}
