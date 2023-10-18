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
	print("\n")
	print(ss)
	print("\n")

	return ss
}
func CheckSession(ctx *fasthttp.RequestCtx) string {
	expTime := time.Now().Local().Add(time.Hour*24*30).UnixNano() / 1000
	// go GenerateJWT([]byte(services.StructToJson(models.AdminCred{AppId: "scm", UserCDB: "WVhCcGEyVjVMWFl5TFRNeWQyNDBOelpwZFRRelp6aHNkbXRuYlhBM2QzZGpjM016YTJkM2RERTRPREkxWlRRMGJYWTFjelYy", PassCDB: "TjJKbU9UazJObVJsWXpZMVlqVmlOMkUxTVRJM1pUQTJOVFUxWkdRNU5UUT0=", HostCDB: "Wm1ZeVlUa3lORE10T1RBelpTMDBaRFZrTFRoaVl6QXRNVE14WlRrME9EZGlaVEF4TFdKc2RXVnRhWGd1WTJ4dmRXUmhiblJ1YjNOeGJHUmlMbUZ3Y0dSdmJXRnBiaTVqYkc5MVpBPT0=", UserRedis: "WkdWbVlYVnNkQT09", PassRedis: "TUdFek9EZzJZMkl3TXpZME5EUm1aV0l3WXpVM01UY3dOV0UyWldKa04yST0=", HostRedis: "WVhCdU1TMXJaWGt0Wm1sdVkyZ3RNelExTnpZdWRYQnpkR0Z6YUM1cGJ3PT0=", PortRedis: "TXpRMU56WT0="})), expTime, ctx)

	go GenerateJWT([]byte(services.StructToJson(models.AdminCred{AppId: "scm", UserCDB: "WVdSdGFXND0=", PassCDB: "TVRJeg==", HostCDB: "Ykc5allXeG9iM04wT2pVNU9EUT0=", HostRedis: "WVhCdU1TMXJaWGt0Wm1sdVkyZ3RNelExTnpZdWRYQnpkR0Z6YUM1cGJ3PT0=", PortRedis: "TXpRMU56WT0="})), expTime, ctx)

	authHeader := ctx.Request.Header.Peek("Authorization")
	tokenString, err := extractBearerToken(authHeader)
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
				var adminCred models.AdminCred
				claims := token.Claims.(jwt.MapClaims)
				data := claims["data"].(string)
				services.JsonToStruct(string(data), &adminCred)
				if adminCred.HostCDB != "" {
					config.CDB_HOST_ADMIN = adminCred.HostCDB
					config.CDB_USER_ADMIN = adminCred.UserCDB
					config.CDB_PASS_ADMIN = adminCred.PassCDB
					print("CDB credential has been set\n\n")
					return "Accessed by Admin"
				}
				if adminCred.HostRedis != "" {
					config.REDIS_HOST = adminCred.HostRedis
					config.REDIS_USER = adminCred.UserRedis
					config.REDIS_PASS = adminCred.PassRedis
					config.REDIS_PORT = adminCred.PortRedis
				}
				sessionData := services.GetValueRedis(claim.IdUser)
				if sessionData == "" {
					services.ShowResponseDefault(ctx, fasthttp.StatusUnauthorized, "Session not found")
					return ""
				} else {
					// print(sessionData)
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
