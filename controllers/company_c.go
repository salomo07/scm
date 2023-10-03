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
	resReg, errReg := services.RegisterCompany([]byte(models.StructToJson(dataCompany)))
	if errReg == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadGateway, errReg)
		return
	} else {
		print("\nqwerty" + resReg + "\n")
		var company models.InsertDocumentResponse
		models.JsonToStruct(resReg, &company)
		// json.Unmarshal([]byte(resReg), &company)
		print("\nzxcv" + company.Id)
		createCompanyDB(company.Id)

		// println("\n\n\n")

		// createCompanyDB()
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
