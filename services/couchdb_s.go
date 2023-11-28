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
func FindDocument(body []byte, dbname string) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBAdmin() + dbname + "/_find"
	return SendToNextServer(urlDB, "POST", body)
}
func InsertDocument(body []byte, dbname string) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBAdmin() + dbname
	return SendToNextServer(urlDB, "POST", body)
}
func InsertBulkDocument(body []byte, dbname string) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBAdmin() + dbname + "/_bulk_docs"
	jsonData := `{"docs":` + string(body) + `}`
	return SendToNextServer(urlDB, "POST", []byte(jsonData))
}
func AddUserDB(idcompany string, body []byte) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBAdmin() + "_users/org.couchdb.user:" + idcompany
	return SendToNextServer(urlDB, "PUT", body)
}
func AddAdminRoleForDB(idcompany string, body []byte) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBAdmin() + idcompany + "/_security"
	return SendToNextServer(urlDB, "PUT", body)
}
func UpdateDocument(_id string, data []byte) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBAdmin() + config.TABLE_CORE_NAME + "/" + _id
	return SendToNextServer(urlDB, "PUT", data)
}

// As Company
func InsertDocumentAsComp(company models.Company, body []byte) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBCompany() + company.IdCompany
	return ToCDBCompany(urlDB, "POST", body)
}
func FindDocumentAsComp(company models.Company, body []byte) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBCompany() + company.IdCompany + "/_find"
	return ToCDBCompany(urlDB, "POST", body)
}
func UpdateDocumentAsComp(company models.Company, _iddoc string, data []byte) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBCompany() + company.IdCompany + "/" + _iddoc
	return ToCDBCompany(urlDB, "PUT", data)
}
func DeleteDocumentAsComp(company models.Company, _iddoc string, data []byte) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBCompany() + company.IdCompany + "/" + _iddoc
	return ToCDBCompany(urlDB, "DELETE", data)
}
