package services

import (
	"scm/config"
)

var url string

func init() {
	url = config.GetCredCDB("", "") + "scm_core"
}
func CreateDB(dbname string) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDB("", "") + dbname
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
func AddUserDB(idcompany string, body []byte) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDB("", "") + "_users/org.couchdb.user:" + idcompany
	return SendToNextServer(urlDB, "PUT", body)
}
func AddAdminRoleForDB(idcompany string, body []byte) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDB("", "") + idcompany + "/_security"
	return SendToNextServer(urlDB, "PUT", body)
}
