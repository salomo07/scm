package controllers

import (
	"scm/config"
	"scm/models"
	"scm/services"

	"github.com/valyala/fasthttp"
)

func AddRoleByAdmin(ctx *fasthttp.RequestCtx) {
	if string(ctx.Request.Body()) == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Request body cant be empty")
		return
	}
	var roleModel models.Role
	models.JsonToStruct(string(ctx.PostBody()), &roleModel)
	err := models.ValidateRequiredFields(roleModel)
	if err == "" {
		roleModel.Table = "role"
		query := `{"selector": {"table":"role","code":"` + roleModel.Code + `","idcompany":"` + roleModel.IdCompany + `"},"limit":1}`
		resFind, errFind, stts := services.FindDocument(query, config.DB_CORE_NAME)
		if len(resFind.Docs) == 0 {
			if errFind == "" {
				resBody, errStr, statuscode := services.InsertDocument(models.StructToJson(roleModel), config.DB_CORE_NAME)
				if resBody != "" {
					var insertModel models.InsertResponse
					models.JsonToStruct(resBody, &insertModel)
					roleModel.Id = insertModel.Id
					services.ShowResponseJson(ctx, statuscode, models.StructToJson(roleModel))
				} else {
					services.ShowResponseDefault(ctx, statuscode, errStr)
				}
			} else {
				services.ShowResponseDefault(ctx, stts, errFind)
			}

		} else {
			models.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "'Code' already taken")
		}

	} else {
		models.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, err)
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
		err := models.ValidateRequiredFields(value)
		if err == "" {
			value.Table = "role"
			roleModelTemp = append(roleModelTemp, value)
		} else {
			return
		}
	}
	resBody, errStr, statuscode := services.InsertBulkDocument(models.StructToJson(roleModelTemp), config.DB_CORE_NAME)
	if resBody != "" {
		services.ShowResponseJson(ctx, statuscode, resBody)
	} else {
		services.ShowResponseDefault(ctx, statuscode, errStr)
	}
}
