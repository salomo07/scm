package routers

import (
	"scm/controllers"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func CompanyRouters(router *fasthttprouter.Router) {
	router.POST("/api/v1/admin/company/create/", func(ctx *fasthttp.RequestCtx) {
		//	Endpoint ini hanya bisa diakses oleh SuperAdmin
		_, dbCred, errMsg := controllers.CheckSession(ctx)
		if errMsg == "" {
			controllers.RegisterCompany(dbCred, ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("/api/v1/admin/company/role/addrole", func(ctx *fasthttp.RequestCtx) {
		_, dbCred, errMsg := controllers.CheckSession(ctx)
		if errMsg == "" {
			controllers.AddRole(dbCred, ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("/api/v1/admin/company/copyinitiatedata", func(ctx *fasthttp.RequestCtx) {
		// controllers.CopyInitiateData(adminCred, ctx)

		ctx.Response.Header.Set("Content-Type", "application/json")
	})
}
