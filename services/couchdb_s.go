package services

import (
	"scm/config"
	"scm/models"
)

// Admin
func CreateDB(dbname string) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBAdmin() + dbname
	var xxx []byte
	return SendToNextServer(urlDB, "PUT", xxx)
}
func CreateIndexPerCompany(dbname string) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBAdmin() + dbname
	return SendToNextServer(urlDB, "POST", []byte(`{"index":{"fields":["table","idcompany"]},"name":"companydata","ddoc":"companydata","type":"json"}`))
}
func FindDocument(adminCred string, body []byte, dbname string) (findRes models.FindResponse, errStr string, statuscode int) {
	urlDB := adminCred + dbname + "/_find"
	println(urlDB)
	res, err, code := SendToNextServer(urlDB, "POST", body)
	JsonToStruct(res, &findRes)
	return findRes, err, code
}
func InsertDocument(adminCred string, body []byte, dbname string) (resBody string, errStr string, statuscode int) {
	urlDB := adminCred + dbname
	return SendToNextServer(urlDB, "POST", body)
}
func InsertBulkDocument(adminCred string, body []byte, dbname string) (resBody string, errStr string, statuscode int) {
	urlDB := adminCred + dbname + "/_bulk_docs"
	jsonData := `{"docs":` + string(body) + `}`
	return SendToNextServer(urlDB, "POST", []byte(jsonData))
}
func AddUserDB(adminCred string, idcompany string, body []byte) (resBody string, errStr string, statuscode int) {
	urlDB := adminCred + "_users/org.couchdb.user:" + idcompany
	return SendToNextServer(urlDB, "PUT", body)
}
func AddAdminRoleForDB(adminCred string, idcompany string, body []byte) (resBody string, errStr string, statuscode int) {
	urlDB := adminCred + idcompany + "/_security"
	return SendToNextServer(urlDB, "PUT", body)
}
func UpdateDocument(adminCred string, _id string, data []byte) (resBody string, errStr string, statuscode int) {
	urlDB := adminCred + config.DB_CORE_NAME + "/" + _id
	return SendToNextServer(urlDB, "PUT", data)
}

// As Company
func InsertDocumentAsComp(company models.Company, body []byte) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBCompany(company.UserCDB, company.PassCDB) + company.IdCompany
	return ToCDBCompany(urlDB, "POST", body)
}
func FindDocumentAsComp(company models.Company, body []byte) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBCompany(company.UserCDB, company.PassCDB) + company.IdCompany + "/_find"
	return ToCDBCompany(urlDB, "POST", body)
}
func UpdateDocumentAsComp(company models.Company, _iddoc string, data []byte) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBCompany(company.UserCDB, company.PassCDB) + company.IdCompany + "/" + _iddoc
	return ToCDBCompany(urlDB, "PUT", data)
}
func DeleteDocumentAsComp(company models.Company, _iddoc string, data []byte) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBCompany(company.UserCDB, company.PassCDB) + company.IdCompany + "/" + _iddoc
	return ToCDBCompany(urlDB, "DELETE", data)
}
