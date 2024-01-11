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
func TestDuplicate(adminCred string, ctx *fasthttp.RequestCtx) {

}
func AddMenu(adminCred string, ctx *fasthttp.RequestCtx) {
	if string(ctx.Request.Body()) == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Request body cant be empty")
		return
	}
	var menuModel models.Menu
	models.JsonToStruct(string(ctx.PostBody()), &menuModel)
	err := models.ValidateRequiredFields(menuModel, ctx)
	if err == "" {
		menuModel.Table = "menu"
		if len(menuModel.Submenu) > 0 {
			for i, _ := range menuModel.Submenu {
				menuModel.Submenu[i].IdSubmenu = i + 1
			}
		}
		res, err, stts := services.InsertDocument(adminCred, models.StructToJson(menuModel), config.DB_CORE_NAME)
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
	res, err, sts := services.FindDocument(config.GetCredCDBAdmin(), query, config.DB_CORE_NAME)
	if err == "" {
		if len(res.Docs) > 0 {
			var accessRes models.AccessMenuUpdate
			models.JsonToStruct(models.StructToJson(res.Docs[0]), &accessRes)
			var accessTemp models.AccessMenuUpdate
			models.JsonToStruct(string(ctx.PostBody()), &accessTemp)
			accessTemp.IdAccess = accessRes.IdAccess
			accessTemp.Rev = accessRes.Rev
			resBody, errRes, stscode := services.UpdateDocument(adminCred, accessRes.IdAccess, models.StructToJson(accessTemp))
			if errRes != "" {
				models.ShowResponseDefault(ctx, stscode, errRes)
			} else {
				services.ShowResponseJson(ctx, stscode, resBody)
			}
		} else {
			resBody, errRes, stscode := services.InsertDocument(adminCred, models.StructToJson(accessModel), "scm_core")
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
func AddUser(adminCred models.AdminDB, urlDB string, ctx *fasthttp.RequestCtx) {
	if string(ctx.Request.Body()) == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Request body cant be empty")
		return
	}
	var userModel models.User
	models.JsonToStruct(string(ctx.PostBody()), &userModel)
	err := models.ValidateRequiredFields(userModel, ctx)
	if err == "" {
		encryptedPass := config.EncodingBcrypt(userModel.Password)
		userModel.Password = encryptedPass
		userModel.Table = "user"

		findUserCoreDB := `{"selector":{"username":"` + userModel.Username + `","table":"user"}}`
		log.Println("\n" + adminCred.UserCDB + "\n")
		resFind, errFind, codeFind := services.FindDocumentAsComp(models.Company{UserCDB: adminCred.UserCDB, PassCDB: adminCred.PassCDB}, findUserCoreDB)
		if errFind == "" {
			services.ShowResponseDefault(ctx, codeFind, errFind)
		} else {
			if len(resFind.Docs) > 0 {
				models.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Username already taken")
			} else {
				InsertUser(urlDB)
				resIns, errIns, codeIns := services.InsertDocumentAsComp(models.Company{UserCDB: adminCred.UserCDB, PassCDB: adminCred.PassCDB}, models.StructToJson(userModel))
				if errIns == "" {
					services.ShowResponseJson(ctx, codeIns, resIns)
				} else {
					models.ShowResponseDefault(ctx, fasthttp.StatusInternalServerError, "Gagal melakukan Insert")
				}
			}
		}
	}
}
func InsertUser(urlDB string) {
	print(urlDB)
}
