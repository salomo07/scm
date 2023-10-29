package controllers

import (
	"log"
	"scm/config"
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
	err := models.ValidateStruct(roleModel, ctx)
	if err == "" {
		roleModel.Table = "role"
		resBody, errStr, statuscode := services.InsertDocument([]byte(models.StructToJson(roleModel)), config.TABLE_CORE_NAME)
		if resBody != "" {
			services.ShowResponseJson(ctx, statuscode, resBody)
		} else {
			services.ShowResponseDefault(ctx, statuscode, errStr)
		}
	}
}
func AddRoleBulk(ctx *fasthttp.RequestCtx) {
	if string(ctx.Request.Body()) == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Request body cant be empty")
		return
	}
	var roleModel []models.Role
	var roleModelTemp []models.Role
	models.JsonToStruct(string(ctx.Request.Body()), &roleModel)
	for _, value := range roleModel {
		if value.Name == "" {
			services.ShowResponseDefault(ctx, fasthttp.StatusBadGateway, "name is mandatory")
			return
		} else if value.IdCompany == "" {
			services.ShowResponseDefault(ctx, fasthttp.StatusBadGateway, "idcompany is mandatory")
			return
		}
		value.Table = "role"
		roleModelTemp = append(roleModelTemp, value)
	}
	log.Println(roleModelTemp)
	resBody, errStr, statuscode := services.InsertBulkDocument([]byte(models.StructToJson(roleModelTemp)), config.TABLE_CORE_NAME)
	if resBody != "" {
		services.ShowResponseJson(ctx, statuscode, resBody)
	} else {
		services.ShowResponseDefault(ctx, statuscode, errStr)
	}
}
