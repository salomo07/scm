package controllers

import (
	"scm/config"
	"scm/consts"
	"scm/models"
	"scm/services"
	"strconv"
	"time"

	"github.com/valyala/fasthttp"
)

func RegisterCompany(ctx *fasthttp.RequestCtx) {
	if string(ctx.Request.Body()) == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Request body cant be empty")
		return
	}
	var companyModel models.Company
	models.JsonToStruct(string(ctx.PostBody()), &companyModel)
	companyModel.IdCompany = "IniIDCompanyDummy"
	err := models.ValidateRequiredFields(companyModel)
	if err == "" {
		query := consts.QueryCompanyAlias(companyModel.Alias, companyModel.AppId)

		existCompany, errFind, statuscode := services.FindDocument(query, config.DB_CORE_NAME)
		if errFind != "" {
			services.ShowResponseDefault(ctx, statuscode, errFind)
		} else if len(existCompany.Docs) > 0 {
			services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, `'Alias' has already been used`)
		} else {
			// Insert document company
			companyModel.Table = "company"
			companyModel.IdCompany = "c_" + strconv.FormatInt(time.Now().UnixNano()/1000, 10)
			if companyModel.LevelMembership == "" {
				companyModel.LevelMembership = "default"
			}

			companyInsertRes, errInsert, statuscode := services.InsertDocument(models.StructToJson(companyModel), config.DB_CORE_NAME)
			if errInsert != "" {
				services.ShowResponseDefault(ctx, statuscode, errInsert)
			} else {
				companyJson := models.StructToJson(companyModel)
				createCompanyDB(ctx, companyModel.IdCompany, companyInsertRes, companyJson)

				//Save temporary company data on Redis

				// log.Println(companyModel.IdCompany, companyJson, (time.Hour * 8).String())
				// go services.SaveValueRedis(companyModel.IdCompany, companyJson, (time.Hour * 8).String())
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

			_, err, statuscode := services.AddUserDB(dbName, models.StructToJson(userDBModel))
			if err != "" {
				services.ShowResponseDefault(ctx, statuscode, err)
			} else {
				jsonBody := consts.BodySecurity(dbName)
				_, err, statuscode := services.AddAdminRoleForDB(dbName, jsonBody)
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

				go services.UpdateDocument(insertDocumentResponse.Id, models.StructToJson(companyMod))
				go CopyInitiateData(nil, dbName)
				return
			}
		}
	}
}
func CopyInitiateData(ctx *fasthttp.RequestCtx, idcompany string) {
	query := consts.QueryInit
	res, err, code := services.FindDocument(query, config.DB_CORE_NAME)
	if err != "" {
		services.ShowResponseDefault(ctx, code, err)
	} else {
		var tempData []any
		if len(res.Docs) == 0 {
			services.ShowResponseDefault(ctx, fasthttp.StatusNotFound, "Data default tidak ditemukan")
		} else {
			for _, value := range res.Docs {
				row := models.RemoveField(value, "_rev")
				tempData = append(tempData, row)
			}
			resInsert, errInsert, codeInsert := services.InsertBulkDocument(models.StructToJson(tempData), idcompany)
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
