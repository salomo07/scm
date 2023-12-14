package controllers

import (
	"log"
	"scm/config"
	"scm/models"
	"scm/services"

	"github.com/valyala/fasthttp"
)

func AddCompany() {

}

func AddMenu(ctx *fasthttp.RequestCtx) {
	if string(ctx.Request.Body()) == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Request body cant be empty")
		return
	}
	var menuModel models.Menu
	models.JsonToStruct(string(ctx.PostBody()), &menuModel)
	err := models.ValidateStruct(menuModel, ctx)
	if err == "" {
		menuModel.Table = "menu"
		if len(menuModel.Submenu) > 0 {
			for i, _ := range menuModel.Submenu {
				menuModel.Submenu[i].IdSubmenu = i + 1
			}
		}
		res, err, stts := services.InsertDocument([]byte(models.StructToJson(menuModel)), config.DB_CORE_NAME)
		if err != "" {
			services.ShowResponseJson(ctx, stts, err)
		} else {
			services.ShowResponseJson(ctx, stts, res)
		}
	}
}
func AddAccess(ctx *fasthttp.RequestCtx) {
	if string(ctx.Request.Body()) == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Request body cant be empty")
		return
	}
	var accessModel models.AccessMenu
	models.JsonToStruct(string(ctx.PostBody()), &accessModel)
	accessModel.Table = "access"

	query := `{"selector":{"table":"access","idcompany":"` + accessModel.IdCompany + `","idrole":"` + accessModel.IdRole + `","idmenu":"` + accessModel.Idmenu + `"},"use_index":"_design/companydata","limit":1}`
	print(query)
	res, err, sts := services.FindDocument([]byte(query), config.DB_CORE_NAME)
	if err == "" {
		if len(res.Docs) > 0 {
			var accessRes models.AccessMenuUpdate
			models.JsonToStruct(models.StructToJson(res.Docs[0]), &accessRes)
			var accessTemp models.AccessMenuUpdate
			models.JsonToStruct(string(ctx.PostBody()), &accessTemp)
			accessTemp.IdAccess = accessRes.IdAccess
			accessTemp.Rev = accessRes.Rev
			resBody, errRes, stscode := services.UpdateDocument(accessRes.IdAccess, []byte(models.StructToJson(accessTemp)))
			if errRes != "" {
				models.ShowResponseDefault(ctx, stscode, errRes)
			} else {
				services.ShowResponseJson(ctx, stscode, resBody)
			}
		} else {
			resBody, errRes, stscode := services.InsertDocument([]byte(models.StructToJson(accessModel)), "scm_core")
			if errRes != "" {
				models.ShowResponseDefault(ctx, stscode, errRes)
			} else {
				services.ShowResponseJson(ctx, stscode, resBody)
			}

		}
	} else {
		models.ShowResponseDefault(ctx, sts, err)
	}
}
func AddUser(ctx *fasthttp.RequestCtx) {
	if string(ctx.Request.Body()) == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Request body cant be empty")
		return
	}
	var userModel models.User
	models.JsonToStruct(string(ctx.PostBody()), &userModel)
	err := models.ValidateStruct(userModel, ctx)
	if err == "" {
		encryptedPass := config.EncodingBcrypt(userModel.Password)
		userModel.Password = encryptedPass
		log.Print(models.Company)
		// services.InsertDocumentAsComp()
	} else {
		models.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, err)
	}
}
