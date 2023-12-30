package routers

import (
	"scm/controllers"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func UserRouters(router *fasthttprouter.Router) {
	router.POST("/api/v1/admin/user/create", func(ctx *fasthttp.RequestCtx) {
		adminCred, urlDB, errMsg := controllers.CheckSession(ctx)
		if errMsg == "" {
			controllers.AddUser(adminCred, urlDB, ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("/api/v1/admin/role/create", func(ctx *fasthttp.RequestCtx) {
		_, dbStringCred, errMsg := controllers.CheckSession(ctx)
		if errMsg == "" {
			controllers.AddRole(dbStringCred, ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("/api/v1/admin/role/createbulk", func(ctx *fasthttp.RequestCtx) {
		_, dbStringCred, errMsg := controllers.CheckSession(ctx)
		if errMsg == "" {
			controllers.AddRoleBulk(dbStringCred, ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("/api/v1/admin/user/createbulk", func(ctx *fasthttp.RequestCtx) {
		adminCred, urlDB, errMsg := controllers.CheckSession(ctx)
		if errMsg == "" {
			controllers.AddUser(adminCred, urlDB, ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
}
