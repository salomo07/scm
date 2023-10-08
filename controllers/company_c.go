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

	models.JsonToStruct(string(ctx.PostBody()), &companyModel)
	jsonBody := `{"selector": {"table":"company","alias":"` + companyModel.Alias + `"}}`
	existCompany, errFind, statuscode := services.FindDocument([]byte(jsonBody))
	if errFind != "" {
		services.ShowResponseDefault(ctx, statuscode, errFind)
	} else if existCompany != "" {
		models.JsonToStruct(existCompany, &findResponseModel)
		if len(findResponseModel.Docs) > 0 {
			services.ShowResponseDefault(ctx, statuscode, `"alias" has already been used`)
		} else {
			// Insert document company
			companyModel.IdCompany = "c_" + strconv.FormatInt(time.Now().UnixNano()/1000, 10)
			_, errInsert, statuscode := services.InsertDocument([]byte(models.StructToJson(companyModel)))
			if errInsert != "" {
				services.ShowResponseDefault(ctx, statuscode, errInsert)
			} else {
				createCompanyDB(ctx, companyModel.IdCompany)
				services.ShowResponseDefault(ctx, statuscode, "Data saved successfully")
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

		}
	}
}
func SetRoleCompanyDB(ctx *fasthttp.RequestCtx) {
	session := CheckSession(ctx)
	println(session)
	if session != "" {

	} else {

	}
}
