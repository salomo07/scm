package routers

import (
	"scm/models"
	"scm/services"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func CouchDBRouters(router *fasthttprouter.Router) {
	router.POST("/api/v1/createdb/:name", func(ctx *fasthttp.RequestCtx) {
		if ctx.UserValue("name") == "" {
			services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "dbname is needed")
		} else {
			services.CreateDB(ctx.UserValue("name").(string))
		}
	})
	router.POST("/api/v1/createuserdb/:name", func(ctx *fasthttp.RequestCtx) {
		if ctx.UserValue("name") == "" {
			services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "username of db is needed")
		} else {

		}
	})
	router.POST("/api/v1/redis/publish", func(ctx *fasthttp.RequestCtx) {
		var pubdata models.PublishRedis
		models.JsonToStruct(string(ctx.PostBody()), &pubdata)
		services.Publish(pubdata.IdCompany, pubdata.Data)
	})
}
