package controllers

import (
	"scm/models"
	"scm/services"

	"github.com/valyala/fasthttp"
)

func AddCompany() {

}

// @Title Menambahkan menu untuk aplikasi
// @Description Get users related to a specific group.
// @Param  models.RequestHeaders
// @Success  200  object  UsersResponse  "UsersResponse JSON"
// @Failure  400  object  ErrorResponse  "ErrorResponse JSON"
// @Resource users
// @Route /api/group/{groupID}/users [get]
func AddMenu(ctx *fasthttp.RequestCtx) {
	if string(ctx.Request.Body()) == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Request body cant be empty")
		return
	}
	var menuModel models.Menu1
	models.JsonToStruct(string(ctx.PostBody()), &menuModel)
	err := models.ValidateStruct(menuModel, ctx)
	if err == "" {
		menuModel.Table = "menu"
		res, err, stts := services.InsertDocument([]byte(models.StructToJson(menuModel)), "scm_core")
		if err != "" {
			services.ShowResponseJson(ctx, stts, err)
		} else {
			services.ShowResponseJson(ctx, stts, res)
		}
	}
}
func AddAccess1(ctx *fasthttp.RequestCtx) {
	if string(ctx.Request.Body()) == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Request body cant be empty")
		return
	}
	var menuModel models.AccessMenu1
	models.JsonToStruct(string(ctx.PostBody()), &menuModel)
	err := models.ValidateStruct(menuModel, ctx)
	if err == "" {
		menuModel.Table = "access1"
		query := `{"selector":{"idcompany":"c_1697456451227950","table":"access1","idmenu1":"5712a9da17e9ce468530be602523f705"},"use_index":"_design/companydata"}`
		res, err, stts := services.FindDocument([]byte(query), "scm_core")
		if err == "" {
			var findRes models.FindResponse
			models.JsonToStruct(res, &findRes)
			if len(findRes.Docs) > 1 {
				print("\nAccess untuk menu ini sudah ada")
			} else if len(findRes.Docs) == 0 {
				print("\nMenu yang anda pilih tidak dikenali")
			} else {
				print("\nNah ini baru insert access baru")
			}
		}
		// res, err, stts := services.InsertDocument([]byte(models.StructToJson(menuModel)), "scm_core")
		if err != "" {
			services.ShowResponseJson(ctx, stts, err)
		} else {
			services.ShowResponseJson(ctx, stts, res)
		}
	}
}
func AddAccess2(ctx *fasthttp.RequestCtx) {
	if string(ctx.Request.Body()) == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Request body cant be empty")
		return
	}
	var menuModel models.AccessMenu1
	models.JsonToStruct(string(ctx.PostBody()), &menuModel)
	err := models.ValidateStruct(menuModel, ctx)
	if err == "" {
		menuModel.Table = "access1"
		res, err, stts := services.InsertDocument([]byte(models.StructToJson(menuModel)), "scm_core")
		if err != "" {
			services.ShowResponseJson(ctx, stts, err)
		} else {
			services.ShowResponseJson(ctx, stts, res)
		}
	}
}

// {"idcompany":""}
