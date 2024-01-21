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
func AddMenuByAdmin(ctx *fasthttp.RequestCtx) {
	if string(ctx.Request.Body()) == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Request body cant be empty")
		return
	}
	var menuModel models.Menu
	models.JsonToStruct(string(ctx.PostBody()), &menuModel)
	err := models.ValidateRequiredFields(menuModel)

	if err == "" {

		// Cek apakah nama menu sudah dipakai
		queryMenu := `{"selector":{"table":"menu","appid":"` + menuModel.AppId + `","name":"` + menuModel.Name + `"}}`
		resFind, errFind, _ := services.FindDocument(queryMenu, config.DB_CORE_NAME)
		if errFind != "" {
			if len(resFind.Docs) > 0 {
				models.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Menu already exist")
			} else {
				json, err, cod := services.GetDocumentById(config.DB_CORE_NAME, menuModel.AppId)
				if err != "" {
					models.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "App is not exist")
				}
				log.Println(json, err, cod)
			}
		} else {
			services.ShowResponseJson(ctx, fasthttp.StatusInternalServerError, errFind)
		}
	} else {
		models.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, err)
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
		pass := userModel.Password
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
				timemicro := time.Now().UnixNano() / 1000
				userModel.Id = "u_" + strconv.FormatInt(timemicro, 10)

				//Cek dulu apakah Company terdaftar
				resJson, errJson, _ := services.GetDocumentById(config.DB_CORE_NAME, userModel.IdCompany)
				var company models.Company
				models.JsonToStruct(resJson, &company)
				if errJson == "" {
					insertUserData(ctx, userModel, pass)
				} else {
					models.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Company is unregistered")
				}
			}
		}
	} else {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, err)
	}
}
func insertUserData(ctx *fasthttp.RequestCtx, userModel models.User, oldpass string) {
	resIns, errIns, codeIns := services.InsertDocument(models.StructToJson(userModel), config.DB_CORE_NAME)

	if errIns == "" {
		var insertResponseModel models.InsertResponse
		models.JsonToStruct(resIns, &insertResponseModel)

		print("\nAddUserOnCompanyData\n")
		go services.InsertDocument(models.StructToJson(userModel), userModel.IdCompany)
		go AddUserOnCompanyData(userModel.IdCompany, userModel.Id)
		userModel.Password = oldpass
		services.ShowResponseJson(ctx, codeIns, models.StructToJson(userModel))
	} else {
		models.ShowResponseDefault(ctx, fasthttp.StatusInternalServerError, "Gagal melakukan Insert")
	}
}
func AddUserOnCompanyData(idcompany string, iduser string) {
	//Add to DB and Redis
	resjson, err, _ := services.GetDocumentById(config.DB_CORE_NAME, idcompany)
	if err != "" {
		print("Error when getting data by ID")
	} else {
		var company models.CompanyUpdate
		services.JsonToStruct(resjson, &company)
		arrUser := company.Users
		arrUser = addIfNotExist(arrUser, iduser)
		company.Users = arrUser
		go services.UpdateDocument(idcompany, models.StructToJson(company))
	}
}

func addIfNotExist(arr []string, newString string) []string {
	// Check if the new string already exists in the array
	for _, existingString := range arr {
		if existingString == newString {
			return arr
		}
	}
	return append(arr, newString)
}
