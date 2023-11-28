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

func AddUser(ctx *fasthttp.RequestCtx) {
	if string(ctx.Request.Body()) == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Request body cant be empty")
		return
	}
	var userModel models.User
	models.JsonToStruct(string(ctx.PostBody()), &userModel)
	err := models.ValidateStruct(userModel, ctx)
	if err == "" {

	}
	log.Println(userModel)
}
func RegisterCompany(ctx *fasthttp.RequestCtx) {
	if string(ctx.Request.Body()) == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Request body cant be empty")
		return
	}
	var companyModel models.Company
	var findResponseModel models.FindResponse

	models.JsonToStruct(string(ctx.PostBody()), &companyModel)
	jsonBody := `{"selector": {"table":"company","alias":"` + companyModel.Alias + `","limit":1}}`
	existCompany, errFind, statuscode := services.FindDocument([]byte(jsonBody), config.TABLE_CORE_NAME)
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
			companyInsertRes, errInsert, statuscode := services.InsertDocument([]byte(models.StructToJson(companyModel)), config.TABLE_CORE_NAME)
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

func CopyInitiateData(ctx *fasthttp.RequestCtx) {
	if string(ctx.Request.Body()) == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Request body cant be empty")
		return
	}
	var companyData struct {
		IdCompany string `json:"idcompany"`
	}
	models.JsonToStruct(string(ctx.PostBody()), &companyData)
	if companyData.IdCompany == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "idcompany cant empty")
	} else {
		query := `{"selector":{"idcompany":"` + companyData.IdCompany + `"},"use_index":"_design/companydata"}`
		res, err, code := services.FindDocument([]byte(query), "scm_core")
		// println(res)
		if err != "" {
			services.ShowResponseDefault(ctx, code, err)
		} else {
			var findRes models.FindResponse
			services.JsonToStruct(res, &findRes)
			var tempData []any
			if len(findRes.Docs) == 0 {
				services.ShowResponseDefault(ctx, fasthttp.StatusNotFound, "Data default tidak ditemukan")
			} else {
				for _, value := range findRes.Docs {
					xxx := models.RemoveField(value, "_rev")
					tempData = append(tempData, xxx)
				}
				resInsert, errInsert, codeInsert := services.InsertBulkDocument([]byte(models.StructToJson(tempData)), companyData.IdCompany)
				if errInsert != "" {
					services.ShowResponseDefault(ctx, codeInsert, errInsert)
				} else {
					services.ShowResponseJson(ctx, codeInsert, resInsert)
				}

			}

		}

	}
}
