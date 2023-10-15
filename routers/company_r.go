package routers

import (
	"scm/controllers"
	"scm/services"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func CompanyRouters(router *fasthttprouter.Router) {
	router.GET("/company/xxx/", func(ctx *fasthttp.RequestCtx) {
		services.ShowResponseDefault(ctx, 200, "Ini router /company/xxx/")
	})
	router.POST("/company/create/", func(ctx *fasthttp.RequestCtx) {
		if controllers.CheckSession(ctx) != "" {
			controllers.RegisterCompany(ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("/company/adduser/", func(ctx *fasthttp.RequestCtx) {
		if controllers.CheckSession(ctx) != "" {
			controllers.AddUser(ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
}
