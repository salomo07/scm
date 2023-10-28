package controllers

import (
	"fmt"
	"os"
	"scm/config"
	"scm/models"
	"scm/services"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/valyala/fasthttp"
)

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

	return ss
}
func CheckSession(ctx *fasthttp.RequestCtx) string {
	expTime := time.Now().Local().Add(time.Hour*24*30).UnixNano() / 1000
	go GenerateJWT([]byte(services.StructToJson(models.AdminCred{AppId: "scm"})), expTime, ctx)

	// go GenerateJWT([]byte(services.StructToJson(models.AdminCred{AppId: "scm", UserCDB: "WVdSdGFXNWtaWFk9", PassCDB:"WTFKM2IwOUhNRlZxZUdKWFRIZDVXRXRTYUUxaVpYQTBZakZNZWtwV1NYQmhZbWxXWjIwMWFHSlFUMlZ4TkZsVFNrSnlRVXM9", HostCDB: "YUhSMGNEb3ZMek0wTGpFeU9TNHlOeTQyTlRvMU9UZzA=", HostRedis: "WVhCdU1TMXJaWGt0Wm1sdVkyZ3RNelExTnpZdWRYQnpkR0Z6YUM1cGJ3PT0=", PortRedis: "TXpRMU56WT0="})), expTime, ctx)

	authHeader := ctx.Request.Header.Peek("Authorization")
	tokenString, err := extractBearerToken(authHeader)
	if err != nil {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, err.Error())
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
				host := os.Getenv("COUCHDB_HOST")
				print(host)
				var adminCred models.AdminCred
				claims := token.Claims.(jwt.MapClaims)
				data := claims["data"].(string)
				services.JsonToStruct(string(data), &adminCred)
				if adminCred.IdCompany != "" {
					config.CDB_HOST_ADMIN = os.Getenv("COUCHDB_HOST")
					config.CDB_USER_ADMIN = adminCred.UserCDB
					config.CDB_PASS_ADMIN = adminCred.PassCDB
					return "Accessed by Company"
				}
				// print(data)
				return data
			} else {
				print("\n" + "Token is valid" + "\n")
				services.ShowResponseDefault(ctx, fasthttp.StatusUnauthorized, "Token is valid")
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
