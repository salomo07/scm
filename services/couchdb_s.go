package services

import (
	"scm/config"
)

func CreateDB(dbname string) (resBody string, errStr string) {
	url := config.GetCredCDB("", "") + dbname
	print("xxxx" + url)
	var xxx []byte
	return SendToNextServer(url, "PUT", xxx)
}
func RegisterCompany(body []byte) (resBody string, errStr string) {
	url := config.GetCredCDB("", "") + "/scm_core"
	return SendToNextServer(url, "POST", body)
}
