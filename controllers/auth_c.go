package controllers

import (
	"fmt"
	"log"
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
		fmt.Fprintf(ctx, StructToJson(models.DefaultResponse{Status: fasthttp.StatusBadRequest, Messege: "GenerateJWT : " + err.Error()}))
	}
	print(ss)
	return ss
}
func CheckSession(ctx *fasthttp.RequestCtx) bool {
	expTime := time.Now().Local().Add(time.Hour * 8).Unix()
	go GenerateJWT([]byte(StructToJson(models.SessionData{IdCompany: "Company-Xerwerwer", AppId: "wms", IdUser: "iduser-234234235", UserCDB: "admin", PassCDB: "123"})), expTime, ctx)
	authHeader := ctx.Request.Header.Peek("Authorization")
	tokenString, err := extractBearerToken(authHeader)
	ctx.Response.SetStatusCode(fasthttp.StatusUnauthorized)
	if err != nil {
		print(err.Error())
		fmt.Fprintf(ctx, StructToJson(models.DefaultResponse{Status: fasthttp.StatusBadRequest, Messege: err.Error()}))
		return false
	} else {
		token, errToken := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.TOKEN_SALT), nil
		})
		if errToken != nil {
			print("\n" + errToken.Error() + "\n")
			fmt.Fprintf(ctx, StructToJson(models.DefaultResponse{Status: fasthttp.StatusUnauthorized, Messege: errToken.Error()}))
			return false
		} else {
			if token.Valid {
				ctx.Response.SetStatusCode(fasthttp.StatusOK)
				ctx.Response.Header.Set("Content-Type", "application/json")
				claim := claimJWT(token)

				print(claim.IdUser)

				sessionData := services.GetValueRedis(claim.IdUser)
				if sessionData == "" {
					fmt.Fprintf(ctx, StructToJson(models.DefaultResponse{Status: fasthttp.StatusUnauthorized, Messege: "Session not found"}))
					return false
				} else {
					return true
				}
				// fmt.Fprintf(ctx, StructToJson(claim))
			} else {
				print("\n" + "Tidak valis coyyy" + "\n")
				fmt.Fprintf(ctx, StructToJson(models.DefaultResponse{Status: fasthttp.StatusUnauthorized, Messege: "Token is valid"}))
				return false
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
	JsonToStruct(string(data), &session)
	return session
}
func generateJWT(json []byte, expiredtime int64) string {
	mySigningKey := []byte(os.Getenv("TOKEN_SALT"))
	type Claims struct {
		Json string `json:"data"`
		jwt.StandardClaims
	}
	claims := Claims{string(json), jwt.StandardClaims{ExpiresAt: expiredtime}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Println("generateJWT error : ", err)
	}
	return ss
}
