package models

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

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
	UserAgent string `json:"useragent"`
	IpClient  string `json:"ipclient"`
}
type SessionFull struct {
	AdminDB AdminDB    `json:"admindb"`
	User    UserInsert `json:"user"`
}
type AdminDB struct {
	UserCDB string `json:"usercdb"`
	PassCDB string `json:"passcdb"`
}

type LoginResponse struct {
	AppId     string `json:"appid"`
	IdCompany string `json:"idcompany"`
	Token     string `json:"token"`
	Expired   string `json:"expired"`
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
		delete(valueMap, fieldName)
		resultValue = valueMap
	} else {
		fmt.Println("The value is not a map[string]interface{}")
	}

	return resultValue
}
func ValidateRequiredFields(data interface{}, ctx *fasthttp.RequestCtx) string {
	value := reflect.ValueOf(data)
	if value.Kind() != reflect.Struct {
		return "Input is not a struct"
	}

	var missingFields []string

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldName := value.Type().Field(i).Name
		validateTag := value.Type().Field(i).Tag.Get("validate")

		if validateTag == "required" {
			fieldValue := field.Interface()
			if fieldValue == "" {
				missingFields = append(missingFields, "'"+fieldName+"'")
			}
		}
	}

	if len(missingFields) > 0 {
		missingFieldsStr := strings.Join(missingFields, ", ")
		if len(missingFields) > 1 {
			ShowResponseDefault(ctx, fasthttp.StatusBadRequest, missingFieldsStr+" fields are required and cannot be empty")
		} else {
			ShowResponseDefault(ctx, fasthttp.StatusBadRequest, missingFieldsStr+" field is required and cannot be empty")
		}
		return missingFieldsStr
	}

	return ""
}
