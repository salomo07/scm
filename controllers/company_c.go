package controllers

import (
	"scm/models"
	"scm/services"
	"strconv"
	"time"

	"github.com/valyala/fasthttp"
)

func AddUser(ctx *fasthttp.RequestCtx) {
	if string(ctx.Request.Body()) == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Request body cant be empty")
		return
	}
	var userModel models.User
	models.JsonToStruct(string(ctx.PostBody()), &userModel)
	services.
}
func RegisterCompany(ctx *fasthttp.RequestCtx) {
	if string(ctx.Request.Body()) == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Request body cant be empty")
		return
	}
	var companyModel models.Company
	var findResponseModel models.FindResponse

	models.JsonToStruct(string(ctx.PostBody()), &companyModel)
	jsonBody := `{"selector": {"table":"company","alias":"` + companyModel.Alias + `"}}`
	existCompany, errFind, statuscode := services.FindDocument([]byte(jsonBody))
	if companyModel.Alias == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "alias is mandatory")
	} else if errFind != "" {
		services.ShowResponseDefault(ctx, statuscode, errFind)
	} else if existCompany != "" {
		models.JsonToStruct(existCompany, &findResponseModel)
		if len(findResponseModel.Docs) > 0 {
			services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, `alias has already been used`)
		} else {
			// Insert document company
			companyModel.IdCompany = "c_" + strconv.FormatInt(time.Now().UnixNano()/1000, 10)
			if companyModel.LevelMembership == "" {
				companyModel.LevelMembership = "default"
			}
			companyInsertRes, errInsert, statuscode := services.InsertDocument([]byte(models.StructToJson(companyModel)))
			if errInsert != "" {
				services.ShowResponseDefault(ctx, statuscode, errInsert)
			} else {
				createCompanyDB(ctx, companyModel.IdCompany, companyInsertRes, models.StructToJson(companyModel))
			}
		}
	}
}

func createCompanyDB(ctx *fasthttp.RequestCtx, dbName string, companyInsertRes string, companyModel string) {
	res, err, statuscode := services.CreateDB(dbName)
	if err != "" {
		services.ShowResponseDefault(ctx, statuscode, err)
	} else {
		var createDBResponse models.CreateDBResponse
		models.JsonToStruct(res, &createDBResponse)
		if createDBResponse.Ok {
			//Tambahkan user dan role untuk DB yang telah dibuat
			var userDBModel models.UserDBModel
			models.JsonToStruct(string(ctx.Request.Body()), &userDBModel)
			userDBModel.Name = dbName
			userDBModel.Password = strconv.FormatInt(time.Now().UnixNano()/1000, 10)
			userDBModel.Type = "user"
			userDBModel.Roles = []string{"admin_role"}

			_, err, statuscode := services.AddUserDB(dbName, []byte(models.StructToJson(userDBModel)))
			if err != "" {
				services.ShowResponseDefault(ctx, statuscode, err)
			} else {
				jsonSecurity := `{"admins": {"names": ["` + dbName + `"],"roles": ["admin_role"]},"members": {"names": [],"roles": []}}`
				_, err, statuscode := services.AddAdminRoleForDB(dbName, []byte(jsonSecurity))
				if err != "" {
					services.ShowResponseDefault(ctx, statuscode, err)
				}
				services.ShowResponseJson(ctx, statuscode, `{"idcompany":"`+dbName+`","usercdb":"`+dbName+`","passcdb":"`+userDBModel.Password+`","messege":"Company was saved"}`)
				var insertDocumentResponse models.InsertDocumentResponse
				models.JsonToStruct(companyInsertRes, &insertDocumentResponse)
				var companyMod models.CompanyEdit
				models.JsonToStruct(companyModel, &companyMod)
				companyMod.UserCDB = dbName
				companyMod.PassCDB = userDBModel.Password
				companyMod.Rev = insertDocumentResponse.Rev

				go services.UpdateDocument(insertDocumentResponse.Id, []byte(models.StructToJson(companyMod)))
				return
			}
		}
	}
}
