package config

import (
	"encoding/base64"
	"os"

	"golang.org/x/crypto/bcrypt"
)

var TOKEN_SALT = "RHJlYW1UaGVhdGVy"

// var CDB_HOST_ADMIN = "10.180.8.74"
// var CDB_HOST_ADMIN = "localhost"
// var CDB_USER_ADMIN = "admin"
// var CDB_PASS_ADMIN = "123"
// var CDB_PORT_ADMIN = "5984"
// var CDB_CRED_ADMIN = ""
var CDB_HOST_ADMIN = "Wm1ZeVlUa3lORE10T1RBelpTMDBaRFZrTFRoaVl6QXRNVE14WlRrME9EZGlaVEF4TFdKc2RXVnRhWGd1WTJ4dmRXUmhiblJ1YjNOeGJHUmlMbUZ3Y0dSdmJXRnBiaTVqYkc5MVpBPT0="
var CDB_USER_ADMIN = "WVhCcGEyVjVMWFl5TFRNeWQyNDBOelpwZFRRelp6aHNkbXRuYlhBM2QzZGpjM016YTJkM2RERTRPREkxWlRRMGJYWTFjelYy"
var CDB_PASS_ADMIN = "TjJKbU9UazJObVJsWXpZMVlqVmlOMkUxTVRJM1pUQTJOVFUxWkdRNU5UUT0="
var CDB_CRED_ADMIN = ""
var isLocal = false

var REDIS_CRED = ""
var REDIS_USER = "WkdWbVlYVnNkQT09"
var REDIS_PASS = "TUdFek9EZzJZMkl3TXpZME5EUm1aV0l3WXpVM01UY3dOV0UyWldKa04yST0="
var REDIS_HOST = "WVhCdU1TMXJaWGt0Wm1sdVkyZ3RNelExTnpZdWRYQnpkR0Z6YUM1cGJ3PT0="
var REDIS_PORT = "TXpRMU56WT0="

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

func GetCredCDB(userdb string, passdb string) string {
	protocol := ""
	if !isLocal {
		protocol = "https://"
		return GetCredCDBFromIBM()
	} else {
		protocol = "http://"
	}
	CDB_CRED_ADMIN = protocol + userdb + ":" + passdb + "@" + CDB_HOST_ADMIN + "/"
	return CDB_CRED_ADMIN
}

func GetCredRedis() string {
	// REDIS_CRED = "redis://localhost:5984" //Local
	if REDIS_CRED != "" {
		return REDIS_CRED
	}
	for x := 0; x < 2; x++ {
		res, err := DecodedCredtial(REDIS_USER)
		if err != "" {
			print(err)
		}
		REDIS_USER = res
	}
	for x := 0; x < 2; x++ {
		res, err := DecodedCredtial(REDIS_PASS)
		if err != "" {
			print(err)
		}
		REDIS_PASS = res
	}
	for x := 0; x < 2; x++ {
		res, err := DecodedCredtial(REDIS_HOST)
		if err != "" {
			print(err)
		}
		REDIS_HOST = res
	}

	for x := 0; x < 2; x++ {
		res, err := DecodedCredtial(REDIS_PORT)
		if err != "" {
			print(err)
		}
		REDIS_PORT = res
	}
	REDIS_CRED = "redis://" + REDIS_USER + ":" + REDIS_PASS + "@" + REDIS_HOST + ":" + REDIS_PORT
	return REDIS_CRED
}

func GetCredCDBFromIBM() string {
	print(CDB_CRED_ADMIN)
	CDB := os.Getenv("CDB_USER_ADMIN")
	if CDB != "" {
		return CDB
	}
	if CDB_CRED_ADMIN != "" {
		return CDB_CRED_ADMIN
	}
	for x := 0; x < 2; x++ {
		res, err := DecodedCredtial(CDB_USER_ADMIN)
		if err != "" {
			print(err)
		}
		CDB_USER_ADMIN = res
	}

	for x := 0; x < 2; x++ {

		res, err := DecodedCredtial(CDB_PASS_ADMIN)
		if err != "" {
			print(err)
		}
		CDB_PASS_ADMIN = res
	}
	for x := 0; x < 2; x++ {
		res, err := DecodedCredtial(CDB_HOST_ADMIN)
		if err != "" {
			print(err)
		}
		CDB_HOST_ADMIN = res
	}
	CDB_CRED_ADMIN = "https://" + CDB_USER_ADMIN + ":" + CDB_PASS_ADMIN + "@" + CDB_HOST_ADMIN + "/"
	print(CDB_CRED_ADMIN)
	return CDB_CRED_ADMIN
}
