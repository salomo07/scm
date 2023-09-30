package routers

import (
	"fmt"
	"scm/controllers"
	"scm/models"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func UsersRouters(router *fasthttprouter.Router) {
	router.POST("/company/create/:idcompany", func(ctx *fasthttp.RequestCtx) {
		if controllers.CheckSession(ctx) {
			controllers.CreateCompany(ctx)
		} else {
			ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
	router.POST("/company/create/", func(ctx *fasthttp.RequestCtx) {
		fmt.Fprintf(ctx, controllers.StructToJson(models.DefaultResponse{Status: fasthttp.StatusBadRequest, Messege: "idcompany is needed"}))
		ctx.Response.Header.Set("Content-Type", "application/json")
	})

	// router.POST("/users/create", SendToNextServer)
	// router.POST("/lb/:appid/*path", SendToNextServer)

	// router.GET("/servers/:appid", GetServerPool)    //Menampilkan pool server berdasarkan appid
	// router.POST("/servers/:appid", AddServerHandle) //Menyimpan pool server berdasarkan appid
}
