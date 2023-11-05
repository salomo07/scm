package routers

import (
	"scm/controllers"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func UserRouters(router *fasthttprouter.Router) {
	router.POST("api/v1/admin/user/create", func(ctx *fasthttp.RequestCtx) {
		if controllers.CheckSession(ctx) != "" {
			controllers.AddUser(ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("api/v1/admin/role/create", func(ctx *fasthttp.RequestCtx) {
		if controllers.CheckSession(ctx) != "" {
			controllers.AddRole(ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("api/v1/admin/role/createbulk", func(ctx *fasthttp.RequestCtx) {
		if controllers.CheckSession(ctx) != "" {
			controllers.AddRoleBulk(ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("api/v1/admin/user/createbulk", func(ctx *fasthttp.RequestCtx) {
		if controllers.CheckSession(ctx) != "" {
			controllers.AddUser(ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
}
