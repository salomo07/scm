package controllers

import (
	"scm/config"
	"scm/models"
	"scm/services"

	"github.com/valyala/fasthttp"
)

func AddRole(adminCred string, ctx *fasthttp.RequestCtx) {
	if string(ctx.Request.Body()) == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Request body cant be empty")
		return
	}
	var roleModel models.Role
	models.JsonToStruct(string(ctx.PostBody()), &roleModel)
	err := models.ValidateStruct(roleModel, ctx)
	if err == "" {
		roleModel.Table = "role"
		resBody, errStr, statuscode := services.InsertDocument(adminCred, []byte(models.StructToJson(roleModel)), config.DB_CORE_NAME)
		if resBody != "" {
			services.ShowResponseJson(ctx, statuscode, resBody)
		} else {
			services.ShowResponseDefault(ctx, statuscode, errStr)
		}
	}
}
func AddRoleBulk(adminCred string, ctx *fasthttp.RequestCtx) {
	if string(ctx.Request.Body()) == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Request body cant be empty")
		return
	}
	var roleModel []models.Role
	var roleModelTemp []models.Role
	models.JsonToStruct(string(ctx.Request.Body()), &roleModel)
	for _, value := range roleModel {
		err := models.ValidateStruct(value, ctx)
		if err == "" {
			value.Table = "role"
			roleModelTemp = append(roleModelTemp, value)
		} else {
			return
		}
	}
	resBody, errStr, statuscode := services.InsertBulkDocument(adminCred, []byte(models.StructToJson(roleModelTemp)), config.DB_CORE_NAME)
	if resBody != "" {
		services.ShowResponseJson(ctx, statuscode, resBody)
	} else {
		services.ShowResponseDefault(ctx, statuscode, errStr)
	}
}
