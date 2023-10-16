package routers

import (
	"scm/controllers"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func CompanyRouters(router *fasthttprouter.Router) {
	router.POST("/admin/company/create/", func(ctx *fasthttp.RequestCtx) {
		if controllers.CheckSession(ctx) != "" {
			controllers.RegisterCompany(ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("/admin/company/adduser/", func(ctx *fasthttp.RequestCtx) {
		if controllers.CheckSession(ctx) != "" {
			controllers.AddUser(ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
}
