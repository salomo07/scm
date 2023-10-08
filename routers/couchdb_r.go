package routers

import (
	"scm/services"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func CouchDBRouters(router *fasthttprouter.Router) {
	router.POST("/createdb/:name", func(ctx *fasthttp.RequestCtx) {
		if ctx.UserValue("name") == "" {
			services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "dbname is needed")
		} else {
			services.CreateDB(ctx.UserValue("name").(string))
		}
		//
	})
	router.POST("/createuserdb/:name", func(ctx *fasthttp.RequestCtx) {
		if ctx.UserValue("name") == "" {
			services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "username of db is needed")
		} else {

		}
		//
	})

}
