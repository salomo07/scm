package config

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

var TOKEN_SALT = "RHJlYW1UaGVhdGVy"

// var CDB_HOST_ADMIN = "10.180.8.74"
// var CDB_HOST_ADMIN = "localhost"
// var CDB_USER_ADMIN = "admin"
// var CDB_PASS_ADMIN = "123"
// var CDB_PORT_ADMIN = "5984"
// var CDB_CRED_ADMIN = ""

var isLocal = false

// var LocalCred = "http://admin:123@10.180.70.66:5984/"
// var LocalCred = "http://admin:123@192.168.0.101:5984/"
var LocalCred = "http://admin:123@1localhost:5984/"

var CDB_USER_ADMIN = "WVdSdGFXND0="
var CDB_PASS_ADMIN = "TVRJeg=="
var CDB_HOST_ADMIN = "TVRreUxqRTJPQzR4TGpRPQ=="
var CDB_CRED_ADMIN = ""

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

func GetCredCDB() string {
	if !isLocal {
		return GetCredCDBFromIBM()
	} else {
		print(LocalCred)
		return LocalCred
	}
}

func GetCredRedis() string {
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
	print()
	return CDB_CRED_ADMIN
}
