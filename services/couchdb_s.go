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
func FindDocument(query string, dbname string) (findRes models.FindResponse, errStr string, statuscode int) {
	urlDB := config.GetCredCDBAdmin() + dbname + "/_find"
	log.Println(urlDB, "POST", query)
	res, err, code := SendToNextServer(urlDB, "POST", query)
	JsonToStruct(res, &findRes)
	if code > 303 {
		return models.FindResponse{}, res, code
	}
	return findRes, err, code
}
func GetDocumentById(dbname string, id string) (resjson string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBAdmin() + dbname + "/" + id
	res, err, code := SendToNextServer(urlDB, "GET", "")
	log.Println(urlDB, "GET", "", res, err, code)
	resjson = res
	if code > 303 {
		return "", resjson, code
	}
	return resjson, err, code
}
func PutDocument(body string, dbname string, iddocument string) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBAdmin() + dbname + "/" + iddocument
	return SendToNextServer(urlDB, "PUT", body)
}
func InsertDocument(body string, dbname string) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBAdmin() + dbname
	// log.Println(urlDB, "POST", body)
	return SendToNextServer(urlDB, "POST", body)
}
func InsertBulkDocument(body string, dbname string) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBAdmin() + dbname + "/_bulk_docs"
	jsonData := `{"docs":` + body + `}`
	return SendToNextServer(urlDB, "POST", jsonData)
}
func AddUserDB(idcompany string, body string) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBAdmin() + "_users/org.couchdb.user:" + idcompany
	return SendToNextServer(urlDB, "PUT", body)
}
func AddAdminRoleForDB(idcompany string, body string) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBAdmin() + idcompany + "/_security"
	return SendToNextServer(urlDB, "PUT", body)
}
func UpdateDocument(_id string, data string) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBAdmin() + config.DB_CORE_NAME + "/" + _id
	log.Println(urlDB, "PUT", data)
	return SendToNextServer(urlDB, "PUT", data)
}

// As Company
func InsertDocumentAsComp(company models.Company, body string) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBCompany(company.UserCDB, company.PassCDB) + company.IdCompany
	return ToCDBCompany(urlDB, "POST", body)
}
func FindDocumentAsComp(company models.Company, query string) (findRes models.FindResponse, errStr string, statuscode int) {
	urlDB := config.GetCredCDBCompany(company.UserCDB, company.PassCDB) + company.IdCompany + "/_find"

	print(urlDB + "\n" + query)
	resBody, err, code := ToCDBCompany(urlDB, "POST", query)
	JsonToStruct(resBody, &findRes)
	if code > 303 && config.UsingIBM == true {
		return models.FindResponse{}, resBody, code
	}
	return findRes, err, code
}
func GetDocumentByIdAsComp(company models.Company, dbname string, id string) (resjson string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBCompany(company.UserCDB, company.PassCDB) + dbname + "/" + id
	res, err, code := SendToNextServer(urlDB, "GET", "")
	resjson = res
	if code > 303 && config.UsingIBM == true {
		return "", resjson, code
	}
	return resjson, err, code
}
func UpdateDocumentAsComp(company models.Company, _iddoc string, data string) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBCompany(company.UserCDB, company.PassCDB) + company.IdCompany + "/" + _iddoc
	return ToCDBCompany(urlDB, "PUT", data)
}
func DeleteDocumentAsComp(company models.Company, _iddoc string, data string) (resBody string, errStr string, statuscode int) {
	urlDB := config.GetCredCDBCompany(company.UserCDB, company.PassCDB) + company.IdCompany + "/" + _iddoc
	return ToCDBCompany(urlDB, "DELETE", data)
}
