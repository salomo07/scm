package routers

import (
	"scm/controllers"
	"scm/services"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func CompanyRouters(router *fasthttprouter.Router) {
	router.POST("/api/v1/admin/company/create/", func(ctx *fasthttp.RequestCtx) {
		adminDB, _, errMsg := controllers.CheckSession(ctx)

		//Endpoint ini hanya bisa diakses oleh SuperAdmin (bukan company)
		if errMsg == "Company is unregistered" && adminDB.UserCDB == "" {
			controllers.RegisterCompany(ctx)
		} else {
			services.ShowResponseDefault(ctx, fasthttp.StatusUnauthorized, "You have not access to this endpoint.")
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("/api/v1/admin/company/role/addrole", func(ctx *fasthttp.RequestCtx) {
		_, _, errMsg := controllers.CheckSession(ctx)
		if errMsg == "" {
			print("AddRoleByAdmin\n")
			controllers.AddRoleByAdmin(ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("/api/v1/admin/company/copyinitiatedata", func(ctx *fasthttp.RequestCtx) {
		// controllers.CopyInitiateData(adminCred, ctx)

		ctx.Response.Header.Set("Content-Type", "application/json")
	})
}
