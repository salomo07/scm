package routers

import (
	"scm/controllers"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func CouchDBRouters(router *fasthttprouter.Router) {
	router.POST("/createdb/:name", func(ctx *fasthttp.RequestCtx) {
		controllers.CreateCompanyDB(ctx)
	})
}
