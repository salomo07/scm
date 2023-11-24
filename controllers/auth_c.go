package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"scm/config"
	"scm/models"
	"scm/services"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/valyala/fasthttp"
)

func Logining(ctx *fasthttp.RequestCtx) string {
	rawJSON := ctx.Request.Body()
	var loginInput models.LoginInput
	err := json.Unmarshal(rawJSON, &loginInput)
	if err == nil {
		log.Println(loginInput)
		jsonResponse, errResponse := services.GetValueRedis(loginInput.Username)
		if errResponse == "" {
			print(services.GetValueRedis("Salomo07"))
			token, errToken := jwt.Parse(jsonResponse, func(token *jwt.Token) (interface{}, error) {
				return []byte(config.TOKEN_SALT), nil
			})
			if errToken != nil {
				services.ShowResponseDefault(ctx, fasthttp.StatusUnauthorized, errToken.Error())
				return ""
			} else {
				if token.Valid {

				} else {
					CheckingToDB()
				}
			}
		}
		print("\n")
	}

	return ""
}
func CheckingToDB() {

}
func GenerateJWT(json []byte, expiredtime int64, ctx *fasthttp.RequestCtx) string {
	mySigningKey := []byte(config.TOKEN_SALT)
	type Claims struct {
		Json string `json:"data"`
		jwt.StandardClaims
	}
	claims := Claims{string(json), jwt.StandardClaims{ExpiresAt: expiredtime}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Fprintf(ctx, services.StructToJson(models.DefaultResponse{Status: fasthttp.StatusBadRequest, Messege: "GenerateJWT : " + err.Error()}))
	}
	print(ss)
	return ss
}
func CheckAdminKey(key string) string {
	val, err := services.GetValueRedis(key)
	if err != "" {
		print("Error : " + err + "\n\n")
		return ""
	} else {
		return val
	}
}
func CheckSession(ctx *fasthttp.RequestCtx) string {
	// expTime := time.Now().Local().Add(time.Hour*24*30).UnixNano() / 1000
	// go GenerateJWT([]byte(services.StructToJson(models.AdminCred{AppId: "scm", AdminKey: "$2a$10$4IKUOc7Y9/ofzqik6B73/unL4EQfGExo.jeObRO5Rt9JQ2Q6qcJxG"})), expTime, ctx)

	// go GenerateJWT([]byte(services.StructToJson(models.AdminCred{AppId: "scm", UserCDB: "WVdSdGFXNWtaWFk9", PassCDB: "WTFKM2IwOUhNRlZxZUdKWFRIZDVXRXRTYUUxaVpYQTBZakZNZWtwV1NYQmhZbWxXWjIwMWFHSlFUMlZ4TkZsVFNrSnlRVXM9"})), expTime, ctx)

	authHeader := ctx.Request.Header.Peek("Authorization")
	tokenString, err := extractBearerToken(authHeader)
	if err != nil {
		services.ShowResponseDefault(ctx, fasthttp.StatusUnauthorized, err.Error())
		return ""
	} else {
		token, errToken := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.TOKEN_SALT), nil
		})

		if errToken != nil {
			services.ShowResponseDefault(ctx, fasthttp.StatusUnauthorized, errToken.Error())
			return ""
		} else {
			if token.Valid {
				ctx.Response.SetStatusCode(fasthttp.StatusOK)

				var adminCred models.AdminCred
				claims := token.Claims.(jwt.MapClaims)
				data := claims["data"].(string)
				services.JsonToStruct(string(data), &adminCred)
				if adminCred.IdCompany != "" {
					config.CDB_HOST_ADMIN = os.Getenv("COUCHDB_HOST")
					config.CDB_USER_ADMIN = os.Getenv("COUCHDB_USER")
					config.CDB_PASS_ADMIN = os.Getenv("COUCHDB_PASSWORD")
					return "Accessed by Company"
				} else if adminCred.AdminKey != "" && adminCred.AdminKey == CheckAdminKey("apikeyscm") {
					print("API key is valid\n" + data + "\n\n")
					return data
				}
				print("API key is invalid\n")
				services.ShowResponseDefault(ctx, fasthttp.StatusUnauthorized, "Token is invalid")
				return ""
			} else {
				print("\n" + "Token is invalid" + "\n")
				services.ShowResponseDefault(ctx, fasthttp.StatusUnauthorized, "Token is invalid")
				return ""
			}
		}
	}
}

func extractBearerToken(authHeader []byte) (string, error) {
	// Check if the Authorization header starts with "Bearer "
	if !strings.HasPrefix(string(authHeader), "Bearer ") {
		return "", fmt.Errorf("invalid Bearer token format")
	}

	// Extract the token (remove "Bearer " prefix)
	token := strings.TrimPrefix(string(authHeader), "Bearer ")

	return token, nil
}
func claimJWT(token *jwt.Token) (session models.SessionData) {
	claims := token.Claims.(jwt.MapClaims)
	data := claims["data"].(string)
	services.JsonToStruct(string(data), &session)
	return session
}
