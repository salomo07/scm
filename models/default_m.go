package models

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
	"gopkg.in/go-playground/validator.v9"
)

type DefaultResponse struct {
	Status  int    `json:"status"`
	Messege string `json:"messege"`
}

type SessionData struct {
	IdCompany string `json:"idcompany" validate:"required"`
	IdUser    string `json:"iduser" validate:"required"`
	AppId     string `json:"appid" validate:"required"`
	UserCDB   string `json:"ucdb" validate:"required"`
	PassCDB   string `json:"pcdb" validate:"required"`
}
type AdminCred struct {
	AppId     string `json:"appid"`
	IdCompany string `json:"idcompany"`
	UserCDB   string `json:"usercdb"`
	PassCDB   string `json:"passcdb"`
	AdminKey  string `json:"adminkey"`
}
type LoginCred struct {
	AppId     string `json:"appid"`
	IdCompany string `json:"idcompany"`
	PassCDB   string `json:"passcdb"`
	UserApp   string `json:"userapp"`
	PassApp   string `json:"passapp"`
}

func JsonToStruct(jsonStr string, dynamic any) interface{} {
	json.Unmarshal([]byte(jsonStr), &dynamic)
	return dynamic
}
func StructToJson(v any) string {
	res, err := json.Marshal(v)
	if err != nil {
		println("Fail to convert to JSON")
	}
	return string(res)
}
func ValidateStruct(myStruct any, ctx *fasthttp.RequestCtx) (err string) {
	validate := validator.New()
	if err := validate.Struct(myStruct); err != nil {
		ShowResponseDefault(ctx, fasthttp.StatusBadRequest, err.Error())
		return err.Error()
	}
	return ""
}
func ShowResponseDefault(ctx *fasthttp.RequestCtx, statuscode int, msg string) {
	ctx.Response.SetStatusCode(statuscode)
	fmt.Fprintf(ctx, StructToJson(DefaultResponse{Status: statuscode, Messege: msg}))
}
