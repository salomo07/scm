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

func AddMenu(adminCred string, ctx *fasthttp.RequestCtx) {
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
		res, err, stts := services.InsertDocument(adminCred, []byte(models.StructToJson(menuModel)), config.DB_CORE_NAME)
		if err != "" {
			services.ShowResponseJson(ctx, stts, err)
		} else {
			services.ShowResponseJson(ctx, stts, res)
		}
	}
}
func AddAccess(adminCred string, ctx *fasthttp.RequestCtx) {
	if string(ctx.Request.Body()) == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Request body cant be empty")
		return
	}
	var accessModel models.AccessMenu
	models.JsonToStruct(string(ctx.PostBody()), &accessModel)
	accessModel.Table = "access"

	query := `{"selector":{"table":"access","idcompany":"` + accessModel.IdCompany + `","idrole":"` + accessModel.IdRole + `","idmenu":"` + accessModel.Idmenu + `"},"use_index":"_design/companydata","limit":1}`
	print(query)
	res, err, sts := services.FindDocument(config.GetCredCDBAdmin(), []byte(query), config.DB_CORE_NAME)
	if err == "" {
		if len(res.Docs) > 0 {
			var accessRes models.AccessMenuUpdate
			models.JsonToStruct(models.StructToJson(res.Docs[0]), &accessRes)
			var accessTemp models.AccessMenuUpdate
			models.JsonToStruct(string(ctx.PostBody()), &accessTemp)
			accessTemp.IdAccess = accessRes.IdAccess
			accessTemp.Rev = accessRes.Rev
			resBody, errRes, stscode := services.UpdateDocument(adminCred, accessRes.IdAccess, []byte(models.StructToJson(accessTemp)))
			if errRes != "" {
				models.ShowResponseDefault(ctx, stscode, errRes)
			} else {
				services.ShowResponseJson(ctx, stscode, resBody)
			}
		} else {
			resBody, errRes, stscode := services.InsertDocument(adminCred, []byte(models.StructToJson(accessModel)), "scm_core")
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
func AddUser(adminCred models.AdminCred, urlDB string, ctx *fasthttp.RequestCtx) {
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
		userModel.Table = "user"
		errC := 0
		errM := ""
		for i := 0; i < 2; i++ {
			resBody, errMsg, code := services.InsertDocumentAsComp(userModel.IdCompany, urlDB, []byte(models.StructToJson(userModel)))
			log.Println(userModel)
			if errMsg == "" {
				services.ShowResponseJson(ctx, code, resBody)
				break
			}
			errC = code
			errM = errMsg
		}
		if errC != 0 {
			services.ShowResponseJson(ctx, errC, errM)
		}
	}
}
