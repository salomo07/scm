package routers

import (
	"scm/controllers"
	"scm/models"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func UserRouters(router *fasthttprouter.Router) {
	router.POST("/api/v1/admin/user/create", func(ctx *fasthttp.RequestCtx) {
		_, _, _, isSuperAdmin := controllers.CheckSession(ctx)
		if isSuperAdmin {
			controllers.AddUserByAdmin(ctx)
		} else {
			models.ShowResponseDefault(ctx, fasthttp.StatusUnauthorized, "You have not access to this endpoint (Company is unregistered).")
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("/api/v1/admin/role/create", func(ctx *fasthttp.RequestCtx) {
		_, _, _, isSuperAdmin := controllers.CheckSession(ctx)
		if isSuperAdmin {
			controllers.AddRoleByAdmin(ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("/api/v1/admin/role/createbulk", func(ctx *fasthttp.RequestCtx) {
		_, _, _, isSuperAdmin := controllers.CheckSession(ctx)
		if isSuperAdmin {
			controllers.AddRoleBulk(ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("/api/v1/admin/user/createbulk", func(ctx *fasthttp.RequestCtx) {
		_, _, _, isSuperAdmin := controllers.CheckSession(ctx)
		if isSuperAdmin {
			controllers.AddUserByAdmin(ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
}
