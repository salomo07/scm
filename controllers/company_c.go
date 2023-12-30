package controllers

import (
	"log"
	"scm/config"
	"scm/consts"
	"scm/models"
	"scm/services"
	"strconv"
	"time"

	"github.com/valyala/fasthttp"
)

func RegisterCompany(adminCred string, ctx *fasthttp.RequestCtx) {
	if string(ctx.Request.Body()) == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Request body cant be empty")
		return
	}
	var companyModel models.Company
	models.JsonToStruct(string(ctx.PostBody()), &companyModel)
	query := consts.QueryCompanyAlias(companyModel.Alias)
	existCompany, errFind, statuscode := services.FindDocument(adminCred, query, config.DB_CORE_NAME)
	if companyModel.Alias == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "alias is mandatory")
	} else if errFind != "" {
		services.ShowResponseDefault(ctx, statuscode, errFind)
	} else if len(existCompany.Docs) > 0 {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, `alias has already been used`)
	} else {
		// Insert document company
		companyModel.IdCompany = "c_" + strconv.FormatInt(time.Now().UnixNano()/1000, 10)
		if companyModel.LevelMembership == "" {
			companyModel.LevelMembership = "default"
		}
		log.Println(companyModel)
		companyInsertRes, errInsert, statuscode := services.InsertDocument(adminCred, models.StructToJson(companyModel), config.DB_CORE_NAME)
		if errInsert != "" {
			services.ShowResponseDefault(ctx, statuscode, errInsert)
		} else {
			createCompanyDB(adminCred, ctx, companyModel.IdCompany, companyInsertRes, models.StructToJson(companyModel))
		}
	}
}

func createCompanyDB(adminCred string, ctx *fasthttp.RequestCtx, dbName string, companyInsertRes string, companyModel string) {
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

			_, err, statuscode := services.AddUserDB(adminCred, dbName, models.StructToJson(userDBModel))
			if err != "" {
				services.ShowResponseDefault(ctx, statuscode, err)
			} else {
				jsonBody := consts.BodySecurity(dbName)
				_, err, statuscode := services.AddAdminRoleForDB(adminCred, dbName, jsonBody)
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

				go services.UpdateDocument(adminCred, insertDocumentResponse.Id, models.StructToJson(companyMod))
				go CopyInitiateData(adminCred, nil, dbName)
				return
			}
		}
	}
}

func CopyInitiateData(adminCred string, ctx *fasthttp.RequestCtx, idcompany string) {
	query := consts.QueryInit
	res, err, code := services.FindDocument(config.GetCredCDBAdmin(), query, config.DB_CORE_NAME)
	if err != "" {
		services.ShowResponseDefault(ctx, code, err)
	} else {
		var tempData []any
		if len(res.Docs) == 0 {
			services.ShowResponseDefault(ctx, fasthttp.StatusNotFound, "Data default tidak ditemukan")
		} else {
			for _, value := range res.Docs {
				xxx := models.RemoveField(value, "_rev")
				tempData = append(tempData, xxx)
			}
			resInsert, errInsert, codeInsert := services.InsertBulkDocument(adminCred, models.StructToJson(tempData), idcompany)
			if errInsert != "" {
				services.ShowResponseDefault(ctx, codeInsert, errInsert)
			} else {
				if ctx != nil {
					services.ShowResponseJson(ctx, codeInsert, resInsert)
				}

			}

		}

	}
}
