package controllers

import (
	"encoding/json"
)

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
