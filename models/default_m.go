package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
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
	UserCDB   string `json:"usercdb"`
	PassCDB   string `json:"passcdb"`
	HostCDB   string `json:"hostcdb"`
	UserRedis string `json:"userredis"`
	PassRedis string `json:"passredis"`
	HostRedis string `json:"hostredis"`
	PortRedis string `json:"portredis"`
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
func validateStruct(s any) error {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		tag := field.Tag.Get("validate")

		if tag == "required" {
			// Check if the field is empty
			if isEmpty(value) {
				return errors.New(fmt.Sprintf("Field '%s' is required but is empty.", field.Name))
			}
		}
	}

	return nil
}

func isEmpty(value interface{}) bool {
	switch reflect.TypeOf(value).Kind() {
	case reflect.String:
		return value.(string) == ""
	default:
		return false
	}
}
