package controllers

import (
	"log"
	"scm/config"
	"scm/models"
	"scm/services"
	"strconv"
	"time"

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
func AddUser(adminCred models.AdminDB, ctx *fasthttp.RequestCtx) {
	if string(ctx.Request.Body()) == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Request body cant be empty")
		return
	}
	var userModel models.User
	models.JsonToStruct(string(ctx.PostBody()), &userModel)
	userModel.Id = "IniIdUserDummy"
	userModel.IdCompany = adminCred.UserCDB
	err := models.ValidateRequiredFields(userModel, ctx)
	if err == "" {
		encryptedPass := config.EncodingBcrypt(userModel.Password)
		userModel.Password = encryptedPass
		userModel.Table = "user"
		findUserCoreDB := `{"selector":{"$or": [{"username":"` + userModel.Username + `","table":"user"},{"contact.email":"` + userModel.Contact.Email + `","table":"user"},{"contact.mobile":"` + userModel.Contact.Mobile + `","table":"user"}]}}`
		// ressss := `{"docs":[],"bookmark": "","warning": ""}`
		// resFind, errFind, codeFind := services.FindDocumentAsComp(models.Company{UserCDB: adminCred.UserCDB, PassCDB: adminCred.PassCDB, IdCompany: adminCred.UserCDB}, findUserCoreDB)
		resFind, errFind, codeFind := services.FindDocumentAsComp(models.Company{UserCDB: adminCred.UserCDB, PassCDB: adminCred.PassCDB, IdCompany: adminCred.UserCDB}, findUserCoreDB)

		// models.JsonToStruct(ressss, &resFind)
		if errFind != "" {
			services.ShowResponseDefault(ctx, codeFind, errFind)
		} else {
			jsonHapus := `{"docs":[{
				"_id": "c_1702276535981680",
				"_rev": "3-53b9187200fb5731c2d840707043a1c1",
				"appid": "scm",
				"name": "Suka Makmur Sejahtera",
				"alias": "sms",
				"levelmembership": "default",
				"table": "company",
				"usercdb": "c_1702276535981680",
				"passcdb": "1702276536716766",
				"contact": [
					{
						"email": "",
						"phone": "",
						"mobile": "085186803737"
					}
				],
				"users": [
					"u_1702209986069954"
				]
			}],"bookmark": "g1AAAABUeJzLYWBgYMpgSmHgKy5JLCrJTq2MT8lPzkzJBYoLJccbmhsYGZmbmRqbWloYmlkYgFRywFTiUJMFAIg-FOk",
			"warning": "No matching index found, create an index to optimize query time."}`
			models.JsonToStruct(jsonHapus, &resFind)
			if len(resFind.Docs) > 0 {
				models.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Username, Email or Mobile Phone already taken")
			} else {
				userModel.Id = "u_" + strconv.FormatInt(time.Now().UnixNano()/1000, 10)
				// resIns, errIns, codeIns := services.InsertDocumentAsComp(models.Company{UserCDB: adminCred.UserCDB, PassCDB: adminCred.PassCDB}, models.StructToJson(userModel))
				resIns, errIns, codeIns := services.InsertDocument(config.GetCredCDBAdmin(), models.StructToJson(userModel), userModel.IdCompany)

				if errIns == "" {
					services.ShowResponseJson(ctx, codeIns, resIns)
					go AddUserOnCompanyData(userModel.IdCompany, userModel.Id)
				} else {
					models.ShowResponseDefault(ctx, fasthttp.StatusInternalServerError, "Gagal melakukan Insert")
				}
			}
		}
	}
}
func AddUserOnCompanyData(idcompany string, iduser string) {
	//Add to DB and Redis
	resjson, err, _ := services.GetDocumentById(config.GetCredCDBAdmin(), "scm_core", idcompany)
	if err != "" {
		print("Error when getting data by ID")
	} else {
		var company models.Company
		services.JsonToStruct(resjson, &company)
		log.Println("Ini harusnya diupdate", company)
	}
}
