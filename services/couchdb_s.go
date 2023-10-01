package services

import (
	"scm/config"
)

func CreateDB(userdb string, passdb string, dbname string) {
	url := config.GetCredCDB(userdb, passdb) + dbname
	var xxx []byte
	SendToNextServer(url, "PUT", xxx)
}
