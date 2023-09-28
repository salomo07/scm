package config

import (
	"encoding/base64"
)

var CDB_HOST = "YUhSMGNEb3ZMMnh2WTJGc2FHOXpkQT09"
var CDB_USER = "WVdSdGFXND0="
var CDB_PASS = "TVRJeg=="
var CDB_PORT = "TlRrNE5BPT0="
var CDB_CRED = ""

func DecodedCredtial(encoded string) (string, string) {
	decodedText, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		panic(err)
	}
	return string(decodedText), ""
}

func GetCredCDB() string {
	print(CDB_CRED)
	if CDB_CRED != "" {
		return CDB_CRED
	}
	for x := 0; x < 2; x++ {
		res, err := DecodedCredtial(CDB_USER)
		if err != "" {
			print(err)
		}
		CDB_USER = res
	}

	for x := 0; x < 2; x++ {

		res, err := DecodedCredtial(CDB_PASS)
		if err != "" {
			print(err)
		}
		CDB_PASS = res
	}
	for x := 0; x < 2; x++ {
		res, err := DecodedCredtial(CDB_HOST)
		if err != "" {
			print(err)
		}
		CDB_HOST = res
	}
	for x := 0; x < 2; x++ {
		res, err := DecodedCredtial(CDB_PORT)
		if err != "" {
			print(err)
		}
		CDB_PORT = res
	}
	CDB_CRED = "https://" + CDB_USER + ":" + CDB_PASS + "@" + CDB_HOST + ":" + CDB_PORT

	return CDB_CRED
}
