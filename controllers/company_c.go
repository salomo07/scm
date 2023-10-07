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
		print(findResponseModel)
		if len(findResponseModel.Docs) > 0 {
			services.ShowResponseDefault(ctx, statuscode, `"alias" has already been used`)
		} else {
			// Insert document company
			companyModel.IdCompany = "c_" + strconv.FormatInt(time.Now().UnixNano()/1000, 10)
			_, errInsert, statuscode := services.InsertDocument([]byte(models.StructToJson(companyModel)))
			if errInsert != "" {
				services.ShowResponseDefault(ctx, statuscode, errInsert)
			} else {
				services.ShowResponseDefault(ctx, statuscode, "Data saved successfully")
			}
		}
	} else {
		print("xxxxxxxxxxxxxx")
	}
}

func createCompanyDB(dbName string) {
	services.CreateDB(dbName)
}
func SetRoleCompanyDB(ctx *fasthttp.RequestCtx) {
	session := CheckSession(ctx)
	println(session)
	if session != "" {

	} else {

	}
}
