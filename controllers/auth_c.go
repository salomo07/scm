package controllers

import (
	"fmt"
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
	print(ss)
	return ss
}
func CheckSession(ctx *fasthttp.RequestCtx) string {
	expTime := time.Now().Local().Add(time.Hour * 8).Unix()
	go GenerateJWT([]byte(services.StructToJson(models.SessionData{IdCompany: "Company-Xerwerwer", AppId: "wms", IdUser: "iduser-234234235", UserCDB: "WVdSdGFXND0=", PassCDB: "TVRJeg=="})), expTime, ctx)
	authHeader := ctx.Request.Header.Peek("Authorization")
	tokenString, err := extractBearerToken(authHeader)
	// ctx.Response.SetStatusCode(fasthttp.StatusUnauthorized)
	if err != nil {
		print("\n" + err.Error() + "\n")
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, err.Error())
		return ""
	} else {
		token, errToken := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.TOKEN_SALT), nil
		})
		if errToken != nil {
			print("\n" + errToken.Error() + "\n")
			services.ShowResponseDefault(ctx, fasthttp.StatusUnauthorized, errToken.Error())
			return ""
		} else {
			if token.Valid {
				ctx.Response.SetStatusCode(fasthttp.StatusOK)
				claim := claimJWT(token)

				// print(claim.IdUser)

				sessionData := services.GetValueRedis(claim.IdUser)
				if sessionData == "" {
					services.ShowResponseDefault(ctx, fasthttp.StatusUnauthorized, "Session not found")
					return ""
				} else {
					return sessionData
				}
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
