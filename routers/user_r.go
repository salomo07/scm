package routers

import (
	"scm/controllers"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func UserRouters(router *fasthttprouter.Router) {
	router.POST("/api/v1/admin/user/create", func(ctx *fasthttp.RequestCtx) {
		adminCred, errMsg := controllers.CheckSession(ctx)
		if errMsg == "" {
			controllers.AddUser(adminCred, ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("/api/v1/admin/role/create", func(ctx *fasthttp.RequestCtx) {
		adminCred, errMsg := controllers.CheckSession(ctx)
		if errMsg == "" {
			controllers.AddRole(adminCred, ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("/api/v1/admin/role/createbulk", func(ctx *fasthttp.RequestCtx) {
		adminCred, errMsg := controllers.CheckSession(ctx)
		if errMsg == "" {
			controllers.AddRoleBulk(adminCred, ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("/api/v1/admin/user/createbulk", func(ctx *fasthttp.RequestCtx) {
		adminCred, errMsg := controllers.CheckSession(ctx)
		if errMsg == "" {
			controllers.AddUser(adminCred, ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
}
