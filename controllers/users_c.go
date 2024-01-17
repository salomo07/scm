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
func AddMenu(ctx *fasthttp.RequestCtx) {
	if string(ctx.Request.Body()) == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Request body cant be empty")
		return
	}
	var menuModel models.Menu
	models.JsonToStruct(string(ctx.PostBody()), &menuModel)
	err := models.ValidateRequiredFields(menuModel)
	if err == "" {
		menuModel.Table = "menu"
		if len(menuModel.Submenu) > 0 {
			for i, _ := range menuModel.Submenu {
				menuModel.Submenu[i].IdSubmenu = i + 1
			}
		}
		res, err, stts := services.InsertDocument(models.StructToJson(menuModel), config.DB_CORE_NAME)
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
	res, err, sts := services.FindDocument(query, config.DB_CORE_NAME)
	if err == "" {
		if len(res.Docs) > 0 {
			var accessRes models.AccessMenuUpdate
			models.JsonToStruct(models.StructToJson(res.Docs[0]), &accessRes)
			var accessTemp models.AccessMenuUpdate
			models.JsonToStruct(string(ctx.PostBody()), &accessTemp)
			accessTemp.IdAccess = accessRes.IdAccess
			accessTemp.Rev = accessRes.Rev
			resBody, errRes, stscode := services.UpdateDocument(accessRes.IdAccess, models.StructToJson(accessTemp))
			if errRes != "" {
				models.ShowResponseDefault(ctx, stscode, errRes)
			} else {
				services.ShowResponseJson(ctx, stscode, resBody)
			}
		} else {
			resBody, errRes, stscode := services.InsertDocument(models.StructToJson(accessModel), "scm_core")
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
func CheckUserIsExist(userModel models.User) (findRes models.FindResponse, errStr string, statuscode int) {
	findUserExist := `{"selector":{"$or": [{"username":"` + userModel.Username + `","table":"user"},{"contact.email":"` + userModel.Contact.Email + `","table":"user"},{"contact.mobile":"` + userModel.Contact.Mobile + `","table":"user"}]}}`

	return services.FindDocument(findUserExist, config.DB_CORE_NAME)
}
func AddUserByAdmin(ctx *fasthttp.RequestCtx) {
	if string(ctx.Request.Body()) == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Request body cant be empty")
		return
	}
	var userModel models.User
	models.JsonToStruct(string(ctx.PostBody()), &userModel)
	userModel.Id = "IniIdUserDummy"
	err := models.ValidateRequiredFields(userModel)
	if err == "" {
		encryptedPass := config.EncodingBcrypt(userModel.Password)
		userModel.Password = encryptedPass
		userModel.Table = "user"
		resFind, errFind, codeFind := CheckUserIsExist(userModel)
		if errFind != "" {
			services.ShowResponseDefault(ctx, codeFind, errFind)
		} else {
			if len(resFind.Docs) > 0 {
				models.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Username, Email or Mobile Phone already taken")
			} else {
				userModel.Id = "u_" + strconv.FormatInt(time.Now().UnixNano()/1000, 10)

				resIns, errIns, codeIns := services.InsertDocument(models.StructToJson(userModel), userModel.IdCompany)

				if errIns == "" {
					services.ShowResponseJson(ctx, codeIns, resIns)
					print("\nAddUserOnCompanyData\n")
					go AddUserOnCompanyData(userModel.IdCompany, userModel.Id)
				} else {
					models.ShowResponseDefault(ctx, fasthttp.StatusInternalServerError, "Gagal melakukan Insert")
				}
			}
		}
	} else {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, err)
	}
}
func AddUserOnCompanyData(idcompany string, iduser string) {
	//Add to DB and Redis
	resjson, err, _ := services.GetDocumentById("scm_core", idcompany)
	if err != "" {
		print("Error when getting data by ID")
	} else {
		var company models.Company
		services.JsonToStruct(resjson, &company)
		log.Println("Ini harusnya diupdate", company)
		// services.UpdateDocument(idcompany,)
	}
}
