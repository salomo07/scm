package routers

import (
	"scm/controllers"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func CompanyRouters(router *fasthttprouter.Router) {
	router.POST("/company/create/", func(ctx *fasthttp.RequestCtx) {
		controllers.RegisterCompany(ctx)
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
}
