package controllers

import (
	"scm/services"

	"github.com/valyala/fasthttp"
)

func CreateCompany(ctx *fasthttp.RequestCtx) {
	services.CreateDB("admin", "123", "_users")
}
