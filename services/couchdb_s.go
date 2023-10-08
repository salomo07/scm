package services

import (
	"scm/config"
)

var url string

func init() {
	url = config.GetCredCDB("", "") + "scm_core"
}
func CreateDB(dbname string) (resBody string, errStr string, statuscode int) {
	urlDB := url + dbname
	var xxx []byte
	return SendToNextServer(urlDB, "PUT", xxx)
}
func FindDocument(body []byte) (resBody string, errStr string, statuscode int) {
	urlDB := url + "/_find"
	return SendToNextServer(urlDB, "POST", body)
}
func InsertDocument(body []byte) (resBody string, errStr string, statuscode int) {
	urlDB := url
	return SendToNextServer(urlDB, "POST", body)
}
