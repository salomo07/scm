package routers

import (
	"scm/controllers"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func AuthRouters(router *fasthttprouter.Router) {
	router.POST("/api/v1/auth/login", func(ctx *fasthttp.RequestCtx) {
		// controllers.CheckSession(ctx)
		if controllers.Login(ctx) != "" {
			print("ssss")
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})
}
