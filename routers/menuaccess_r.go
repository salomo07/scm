package routers

import (
	"scm/controllers"
	"scm/models"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func Access_MenuRouters(router *fasthttprouter.Router) {
	router.POST("/api/v1/admin/menu/add", func(ctx *fasthttp.RequestCtx) {
		_, _, _, isSuperAdmin := controllers.CheckSession(ctx)
		if isSuperAdmin {
			return
			controllers.AddMenuByAdmin(ctx)
		} else {
			models.ShowResponseDefault(ctx, fasthttp.StatusUnauthorized, "You have not access to this endpoint.")
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("/api/v1/admin/access1/add", func(ctx *fasthttp.RequestCtx) {
		_, _, _, isSuperAdmin := controllers.CheckSession(ctx)
		if isSuperAdmin {
			controllers.AddAccess(ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
}
