package services

import (
	"scm/config"
)

var url string

func init() {
	url = config.GetCredCDB("", "") + "/scm_core"
}
func CreateDB(dbname string) (resBody string, errStr string) {
	url := config.GetCredCDB("", "") + dbname
	var xxx []byte
	return SendToNextServer(url, "PUT", xxx)
}
func FindDocument(body []byte) (resBody string, errStr string) {
	return SendToNextServer(url, "POST", body)
}
func InsertDocument(body []byte) (resBody string, errStr string) {
	return SendToNextServer(url, "POST", body)
}
