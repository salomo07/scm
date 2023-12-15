package routers

import (
	"scm/controllers"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func CompanyRouters(router *fasthttprouter.Router) {
	router.POST("/api/v1/admin/company/create/", func(ctx *fasthttp.RequestCtx) {
		//	Endpoint ini hanya bisa diakses oleh SuperAdmin
		adminCred, errMsg := controllers.CheckSession(ctx)
		if errMsg == "" {
			controllers.RegisterCompany(adminCred, ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("/api/v1/admin/company/role/addrole", func(ctx *fasthttp.RequestCtx) {
		adminCred, errMsg := controllers.CheckSession(ctx)
		if errMsg == "" {
			controllers.AddRole(adminCred, ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("/api/v1/admin/company/copyinitiatedata", func(ctx *fasthttp.RequestCtx) {
		// if controllers.CheckSession(ctx) != "" {
		// controllers.CopyInitiateData(ctx)
		// }
		adminCred, errMsg := controllers.CheckSession(ctx)
		if errMsg == "" {
			println(adminCred)
			// controllers.CopyInitiateData(adminCred, ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
}
