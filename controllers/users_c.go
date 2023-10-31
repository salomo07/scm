package controllers

import (
	"log"
	"scm/models"
	"scm/services"

	"github.com/valyala/fasthttp"
)

func AddCompany() {

}
func AddMenu(ctx *fasthttp.RequestCtx) {
	if string(ctx.Request.Body()) == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Request body cant be empty")
		return
	}
	var menuModel models.Menu1
	models.JsonToStruct(string(ctx.PostBody()), &menuModel)
	err := models.ValidateStruct(menuModel, ctx)
	if err == "" {

	}
	log.Println(menuModel)
	// services.InsertDocument()
}

// {"idcompany":""}
