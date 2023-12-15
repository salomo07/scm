package routers

import (
	"scm/controllers"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func Access_MenuRouters(router *fasthttprouter.Router) {
	router.POST("/api/v1/admin/menu/add", func(ctx *fasthttp.RequestCtx) {
		adminCred, errMsg := controllers.CheckSession(ctx)
		if errMsg == "" {
			controllers.AddMenu(adminCred, ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("/api/v1/admin/access1/add", func(ctx *fasthttp.RequestCtx) {
		adminCred, errMsg := controllers.CheckSession(ctx)
		if errMsg == "" {
			controllers.AddAccess(adminCred, ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
}
