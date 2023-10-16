package controllers

import (
	"scm/models"
	"scm/services"

	"github.com/valyala/fasthttp"
)

func AddRole(ctx *fasthttp.RequestCtx) {
	if string(ctx.Request.Body()) == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Request body cant be empty")
		return
	}
	var roleModel models.Role
	models.JsonToStruct(string(ctx.PostBody()), &roleModel)
}
