package routers

import (
	"scm/controllers"
	"scm/models"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func UserRouters(router *fasthttprouter.Router) {
	router.POST("/api/v1/admin/user/create", func(ctx *fasthttp.RequestCtx) {
		_, _, errMsg := controllers.CheckSession(ctx)
		if errMsg == "" {
			controllers.AddUserByAdmin(ctx)
		} else {
			models.ShowResponseDefault(ctx, fasthttp.StatusUnauthorized, "You are Unauthorized")
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("/api/v1/admin/role/create", func(ctx *fasthttp.RequestCtx) {
		_, _, errMsg := controllers.CheckSession(ctx)
		if errMsg == "" {
			controllers.AddRole(ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("/api/v1/admin/role/createbulk", func(ctx *fasthttp.RequestCtx) {
		_, _, errMsg := controllers.CheckSession(ctx)
		if errMsg == "" {
			controllers.AddRoleBulk(ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("/api/v1/admin/user/createbulk", func(ctx *fasthttp.RequestCtx) {
		_, _, errMsg := controllers.CheckSession(ctx)
		if errMsg == "" {
			controllers.AddUserByAdmin(ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
}
