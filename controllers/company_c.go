package controllers

import (
	"scm/models"
	"scm/services"
	"strconv"
	"time"

	"github.com/valyala/fasthttp"
)

func RegisterCompany(ctx *fasthttp.RequestCtx) {
	var companyModel models.Company
	var findResponseModel models.FindResponse

	if string(ctx.Request.Body()) == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Request body cant be empty")
		return
	}

	models.JsonToStruct(string(ctx.PostBody()), &companyModel)
	jsonBody := `{"selector": {"table":"company","alias":"` + companyModel.Alias + `"}}`
	existCompany, errFind, statuscode := services.FindDocument([]byte(jsonBody))
	if errFind != "" {
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
			_, errInsert, statuscode := services.InsertDocument([]byte(models.StructToJson(companyModel)))
			if errInsert != "" {
				services.ShowResponseDefault(ctx, statuscode, errInsert)
			} else {
				createCompanyDB(ctx, companyModel.IdCompany)
				services.ShowResponseJson(ctx, statuscode, `{"idcompany":"`+companyModel.IdCompany+`","messege":"Company was saved"}`)
			}
		}
	}
}

func createCompanyDB(ctx *fasthttp.RequestCtx, dbName string) {
	res, err, statuscode := services.CreateDB(dbName)
	if err != "" {
		services.ShowResponseDefault(ctx, statuscode, err)
	} else {
		var createDBResponse models.CreateDBResponse
		models.JsonToStruct(res, &createDBResponse)
		if createDBResponse.Ok {
			//Tambahkan user dan role untuk DB yang telah dibuat
			print("Tambahkan user dan role untuk DB yang telah dibuat")
			var userDBModel models.UserDBModel
			models.JsonToStruct(string(ctx.Request.Body()), &userDBModel)
			// hashPass, _ := HashPassword(dbName)
			userDBModel.Name = dbName
			userDBModel.Password = strconv.FormatInt(time.Now().UnixNano()/1000, 10)
			userDBModel.Type = "user"

			// print(models.StructToJson(userDBModel))
			_, err, statuscode := services.AddUserDB(dbName, []byte(models.StructToJson(userDBModel)))
			if err != "" {
				services.ShowResponseDefault(ctx, statuscode, err)
			} else {
				jsonSecurity := `{"admins": {"names": ["` + dbName + `"], "roles": []}, "members": {"names": [], "roles": []}}`
				setRoleCompanyDB(dbName, []byte(jsonSecurity))
			}
		}
	}
}
func setRoleCompanyDB(dbname string, body []byte) {
	services.AddAdminRoleForDB(dbname, body)
}
