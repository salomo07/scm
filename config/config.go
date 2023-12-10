package config

import (
	"encoding/base64"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var TOKEN_SALT = "RHJlYW1UaGVhdGVy"
var usingIBM = false

var DB_CORE_NAME = "scm_core"
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
	GetCredCDBAdmin()
}
func CompareHashAndPassword(oripass string, hashedPassword string) string {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(oripass))
	if err == nil {
		log.Println("Password is correct!")
		return oripass
	} else if err == bcrypt.ErrMismatchedHashAndPassword {
		log.Println("Password is incorrect.")
		return ""
	} else {
		log.Println("An error occurred:", err)
		return ""
	}
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
	if CDB_CRED_ADMIN != "" {
		return CDB_CRED_ADMIN
	}
	userIBM := os.Getenv("COUCHDB_USER_IBM")
	passIBM := os.Getenv("COUCHDB_PASSWORD_IBM")
	hostIBM := os.Getenv("COUCHDB_HOST_IBM")
	user := os.Getenv("COUCHDB_USER")
	pass := os.Getenv("COUCHDB_PASSWORD")
	host := os.Getenv("COUCHDB_HOST")

	if usingIBM {
		CDB_CRED_ADMIN = "https://" + userIBM + ":" + passIBM + "@" + hostIBM
	} else {
		CDB_CRED_ADMIN = "http://" + user + ":" + pass + "@" + host + ":5984/"
	}
	print(CDB_CRED_ADMIN + "\n")
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
