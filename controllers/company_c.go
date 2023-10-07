package controllers

import (
	"scm/models"
	"scm/services"
	"strconv"
	"time"

	"github.com/valyala/fasthttp"
)

func RegisterCompany(ctx *fasthttp.RequestCtx) {
	var dataCompany models.Company
	dataCompany.IdCompany = "c_" + strconv.FormatInt(time.Now().UnixNano()/1000, 10)
	models.JsonToStruct(string(ctx.PostBody()), &dataCompany)
	// jsonBody := `{"selector": {"table":"company","alias":"`+dataCompany.Alias+`"}}`
	// existCompany, _:=services.FindDocument([]byte(jsonBody))

	resReg, errReg := services.InsertDocument([]byte(models.StructToJson(dataCompany)))
	print("\n\nIni response" + resReg + "\n\n")
	if errReg != "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadGateway, errReg)
		return
	} else {
		var company models.InsertDocumentResponse
		models.JsonToStruct(resReg, &company)
		createCompanyDB(company.Id)
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
