package routers

import (
	"fmt"
	"scm/controllers"
	"scm/models"
	"scm/services"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func CompanyRouters(router *fasthttprouter.Router) {
	router.POST("/company/create/:idcompany", func(ctx *fasthttp.RequestCtx) {
		// controllers.CreateCompanyDB(ctx)
		controllers.RegisterCompany(ctx)
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("/company/create/", func(ctx *fasthttp.RequestCtx) {
		fmt.Fprintf(ctx, services.StructToJson(models.DefaultResponse{Status: fasthttp.StatusBadRequest, Messege: "idcompany is needed"}))
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
}
