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
	CREDADMIN string `json:"credadmin"`
}

type LoginResponse struct {
	AppId     string `json:"appid"`
	UserData  User   `json:"userdata"`
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

type DynamicStruct map[string]interface{}

func RemoveField(original any, fieldName string) DynamicStruct {
	var resultValue DynamicStruct
	if valueMap, ok := original.(map[string]interface{}); ok {
		// Now valueMap is a map[string]interface{}
		fmt.Println(valueMap)
		// models.RemoveField(valueMap, "_rev")
		// services.InsertBulkDocument([]byte(models.StructToJson(findRes.Docs)), companyData.IdCompany)
		delete(valueMap, fieldName)
		resultValue = valueMap
	} else {
		fmt.Println("The value is not a map[string]interface{}")
	}

	return resultValue
}
