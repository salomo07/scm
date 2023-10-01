package controllers

import (
	"scm/models"
	"scm/services"
	"strconv"
	"time"

	"github.com/valyala/fasthttp"
)

func RegisterCompany() {

}
func CreateCompanyDB(ctx *fasthttp.RequestCtx) {
	session := CheckSession(ctx)
	if session != "" {
		var company models.SessionData
		JsonToStruct(session, &company)
		services.CreateDB(company.UserCDB, company.PassCDB, "c_"+strconv.Itoa(time.Now().Nanosecond()))
	}

}
