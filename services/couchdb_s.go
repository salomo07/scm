package services

import (
	"log"
	"scm/config"
	"scm/models"
)

// Admin
func CreateDB(dbname string) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBAdmin() + dbname
	return SendToNextServer(urlDB, "PUT", "")
}
func CreateIndexPerCompany(dbname string) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBAdmin() + dbname
	return SendToNextServer(urlDB, "POST", `{"index":{"fields":["table","idcompany"]},"name":"companydata","ddoc":"companydata","type":"json"}`)
}
func FindDocument(adminCred string, query string, dbname string) (findRes models.FindResponse, errStr string, statuscode int) {
	urlDB := adminCred + dbname + "/_find"
	log.Println(urlDB, "POST", query)
	res, err, code := SendToNextServer(urlDB, "POST", query)
	JsonToStruct(res, &findRes)
	return findRes, err, code
}
func PutDocument(adminCred string, body string, dbname string, iddocument string) (resBody string, errStr string, statuscode int) {
	urlDB := adminCred + dbname + "/" + iddocument
	return SendToNextServer(urlDB, "PUT", body)
}
func InsertDocument(adminCred string, body string, dbname string) (resBody string, errStr string, statuscode int) {
	urlDB := adminCred + dbname
	return SendToNextServer(urlDB, "POST", body)
}
func InsertBulkDocument(adminCred string, body string, dbname string) (resBody string, errStr string, statuscode int) {
	urlDB := adminCred + dbname + "/_bulk_docs"
	jsonData := `{"docs":` + body + `}`
	return SendToNextServer(urlDB, "POST", jsonData)
}
func AddUserDB(adminCred string, idcompany string, body string) (resBody string, errStr string, statuscode int) {
	urlDB := adminCred + "_users/org.couchdb.user:" + idcompany
	return SendToNextServer(urlDB, "PUT", body)
}
func AddAdminRoleForDB(adminCred string, idcompany string, body string) (resBody string, errStr string, statuscode int) {
	urlDB := adminCred + idcompany + "/_security"
	return SendToNextServer(urlDB, "PUT", body)
}
func UpdateDocument(adminCred string, _id string, data string) (resBody string, errStr string, statuscode int) {
	urlDB := adminCred + config.DB_CORE_NAME + "/" + _id
	return SendToNextServer(urlDB, "PUT", data)
}

// As Company
func InsertDocumentAsComp(company models.Company, body string) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBCompany(company.UserCDB, company.PassCDB) + company.IdCompany
	return ToCDBCompany(urlDB, "POST", []byte(body))
}
func FindDocumentAsComp(company models.Company, body string) (findRes models.FindResponse, errStr string, statuscode int) {
	urlDB := config.GetCredCDBCompany(company.UserCDB, company.PassCDB) + company.IdCompany + "/_find"

	resBody, err, code := ToCDBCompany(urlDB, "POST", []byte(body))
	JsonToStruct(resBody, &findRes)
	return findRes, err, code
}
func UpdateDocumentAsComp(company models.Company, _iddoc string, data string) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBCompany(company.UserCDB, company.PassCDB) + company.IdCompany + "/" + _iddoc
	return ToCDBCompany(urlDB, "PUT", []byte(data))
}
func DeleteDocumentAsComp(company models.Company, _iddoc string, data []byte) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBCompany(company.UserCDB, company.PassCDB) + company.IdCompany + "/" + _iddoc
	return ToCDBCompany(urlDB, "DELETE", data)
}
