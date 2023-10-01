package controllers

import (
	"scm/services"

	"github.com/valyala/fasthttp"
)

func CreateCompany(_ *fasthttp.RequestCtx) {
	services.CreateDB("admin", "123", "_users")
}
