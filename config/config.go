package config

import (
	"encoding/base64"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var TOKEN_SALT = "RHJlYW1UaGVhdGVy"
var usingIBM = true

var DB_CORE_NAME = "scm_core"
var CDB_USER_ADMIN = ""
var CDB_PASS_ADMIN = ""

// var CDB_HOST_ADMIN = ""
var CDB_CRED_ADMIN = ""
var CDB_HOST = "192.168.0.102"
var API_KEY_ADMIN = ""
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

	CDB_HOST = os.Getenv("COUCHDB_HOST")

	if usingIBM {
		CDB_CRED_ADMIN = "https://" + userIBM + ":" + passIBM + "@" + hostIBM
	} else {
		CDB_CRED_ADMIN = "http://" + user + ":" + pass + "@" + host + ":5984/"
	}
	print(CDB_CRED_ADMIN + "\n")
	return CDB_CRED_ADMIN
}
func GetCredCDBCompany(user string, pass string) string {
	for i := 0; i < 2; i++ {
		userDec, errUser := DecodedCredtial(user)
		if errUser == "" {
			user = userDec
		}
		passDec, errPass := DecodedCredtial(pass)
		if errPass == "" {
			pass = passDec
		}
	}
	if usingIBM {
		return "https://" + user + ":" + pass + "@" + CDB_HOST
	}
	return "http://" + user + ":" + pass + "@" + CDB_HOST + ":5984/"
}
func GetCredRedis() string {
	return os.Getenv("REDIS_CRED_ADMIN")
}
