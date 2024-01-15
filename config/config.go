package config

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var TOKEN_SALT = "RHJlYW1UaGVhdGVy"
var UsingIBM = true

var DB_CORE_NAME = "scm_core"
var CDB_USER_ADMIN = ""
var CDB_PASS_ADMIN = ""

var CDB_CRED_ADMIN = ""

// var CDB_HOST = "192.168.0.101"
var CDB_HOST = "localhost"
var API_KEY_ADMIN = ""
var REDIS_CRED_ADMIN = ""

func init() {
	er := godotenv.Load()
	if er != nil {
		panic("Fail to load .env file")
	}
	GetCredCDBAdmin()
}
func CompareHashAndPasswordBcrypt(oripass string, hashedPassword string) string {
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

	if UsingIBM {
		CDB_CRED_ADMIN = "https://" + userIBM + ":" + passIBM + "@" + hostIBM
	} else {
		CDB_CRED_ADMIN = "http://" + user + ":" + pass + "@" + host + ":5984/"
	}
	print("Admin DB : " + CDB_CRED_ADMIN + "\n")
	return CDB_CRED_ADMIN
}
func GetCredCDBCompany(user string, pass string) string {
	if UsingIBM {
		return "https://" + user + ":" + pass + "@" + CDB_HOST
	}
	return "http://" + user + ":" + pass + "@" + CDB_HOST + ":5984/"
}
func GetCredRedis() string {
	return os.Getenv("REDIS_CRED_ADMIN")
}

func EncryptAES(plaintext string) (string, error) {
	key := []byte(os.Getenv("KeyEncryptDecrypt"))[:32]
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Printf(err.Error())
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ciphertext := aesGCM.Seal(nil, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(append(nonce, ciphertext...)), nil
}
func DecryptAES(ciphertext string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(os.Getenv("KeyEncryptDecrypt")))
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(data) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := data[:nonceSize], string(data[nonceSize:])
	plaintext, err := aesGCM.Open(nil, nonce, []byte(ciphertext), nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
