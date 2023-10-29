package config

import (
	"encoding/base64"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var TOKEN_SALT = "RHJlYW1UaGVhdGVy"

// var CDB_HOST_ADMIN = "10.180.8.74"
// var CDB_HOST_ADMIN = "localhost"
// var CDB_USER_ADMIN = "admin"
// var CDB_PASS_ADMIN = "123"
// var CDB_PORT_ADMIN = "5984"
// var CDB_CRED_ADMIN = ""

var isLocal = true

// var LocalCred = "http://admin:123@10.180.70.66:5984/"
// var LocalCred = "http://admin:123@192.168.0.101:5984/"
// var LocalCred = "http://admin:123@localhost:5984/"

var TABLE_CORE_NAME = "scm_core"
var CDB_USER_ADMIN = ""
var CDB_PASS_ADMIN = ""
var CDB_HOST_ADMIN = ""
var CDB_CRED_ADMIN = ""

var REDIS_CRED_ADMIN = ""

func init() {
	er := godotenv.Load()
	if er != nil {
		panic("Fail to load .env file")
	}
	user := os.Getenv("COUCHDB_USER")
	pass := os.Getenv("COUCHDB_PASSWORD")
	host := os.Getenv("COUCHDB_HOST")
	CDB_CRED_ADMIN = "http://" + user + ":" + pass + "@" + host + ":5984/"
}
func HashingBcrypt(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		println("Error : ", err)
		return ""
	}
	return string(hashedPassword)
}
func EncodingBcrypt(p string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), 1)
	if err != nil {
		println("Error : ", err)
		return ""
	}
	return string(bytes)
}

func DecodedCredtial(encoded string) (string, string) {
	decodedText, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		panic(err)
	}
	return string(decodedText), ""
}

func GetCredCDBAdmin() string {
	user := os.Getenv("COUCHDB_USER")
	pass := os.Getenv("COUCHDB_PASSWORD")
	host := os.Getenv("COUCHDB_HOST")
	if CDB_CRED_ADMIN == "" {
		CDB_CRED_ADMIN = "http://" + user + ":" + pass + "@" + host + ":5984/"
	}
	print("Admin : " + CDB_CRED_ADMIN)
	return CDB_CRED_ADMIN
}
func GetCredCDBCompany() string {
	if CDB_CRED_ADMIN != "" {
		return CDB_CRED_ADMIN
	}
	CDB_CRED_ADMIN = "http://" + CDB_USER_ADMIN + ":" + CDB_PASS_ADMIN + "@" + CDB_HOST_ADMIN + "/"
	return CDB_CRED_ADMIN
}
func GetCredRedis() string {
	return os.Getenv("REDIS_CRED_ADMIN")
}
