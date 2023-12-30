package models

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/valyala/fasthttp"
)

type DefaultResponse struct {
	Status  int    `json:"status"`
	Messege string `json:"messege"`
}

type SessionToken struct {
	//contoh format IdAppCompanyUser scm*c_1324353452345*u_34345345
	KeyRedis  string `json:"keyredis" validate:"required"`
	AdminKey  string `json:"adminkey"`
	AppId     string `json:"appid"`
	IdCompany string `json:"idcompany"`
	IdUser    string `json:"iduser"`
}
type AdminDB struct {
	UserCDB string `json:"usercdb"`
	PassCDB string `json:"passcdb"`
}

type LoginResponse struct {
	AppId     string     `json:"appid"`
	UserData  UserInsert `json:"userdata"`
	IdCompany string     `json:"idcompany"`
	PassCDB   string     `json:"passcdb"`
	UserApp   string     `json:"userapp"`
	PassApp   string     `json:"passapp"`
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
	errMsg := validate.Struct(myStruct)
	if errMsg != nil {
		// Handle kesalahan validasi
		validationErrors := errMsg.(validator.ValidationErrors)
		for i, e := range validationErrors {
			err = err + e.Namespace() + " is " + e.Tag()
			if i < len(validationErrors)-1 {
				err = err + ", "
			}
			if i == len(validationErrors)-1 {
				err = err + "."
			}
		}
		ShowResponseDefault(ctx, fasthttp.StatusBadRequest, err)
		return err
	}
	return err
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
