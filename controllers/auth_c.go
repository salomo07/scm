package controllers

import (
	"fmt"
	"log"
	"os"
	"scm/config"
	"scm/models"
	"scm/services"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/valyala/fasthttp"
)

func Login(ctx *fasthttp.RequestCtx) string {
	rawJSON := ctx.Request.Body()
	var loginInput models.LoginInput
	models.JsonToStruct(string(rawJSON), &loginInput)
	log.Println(loginInput)
	if loginInput.IdCompany == "" || loginInput.Username == "" || loginInput.Password == "" || loginInput.AppId == "" {
		services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Format input login tidak sesuai")
	} else {
		companyData, err := services.GetValueRedis(loginInput.IdCompany)
		var company models.Company
		services.JsonToStruct(companyData, &company)
		// Jika cred tidak ditemukan di Redis, maka ambil credential dari DB SCM_CRED
		if err != "" {
			services.ShowResponseJson(ctx, fasthttp.StatusInternalServerError, err)
		} else if companyData == "" {
			// services.ShowResponseDefault(ctx, fasthttp.StatusUnauthorized, "Company tidak ditemukan")
			queryFindCred := `{"selectorx":{"_id":"` + loginInput.IdCompany + `","appid":"` + loginInput.AppId + `"}}`
			res, err, code := services.FindDocument(config.GetCredCDBAdmin(), queryFindCred, config.DB_CRED_NAME)
			if err != "" {
				models.ShowResponseDefault(ctx, fasthttp.StatusInternalServerError, err+"Error code : "+strconv.Itoa(code))
			} else {
				if len(res.Docs) == 0 {
					services.ShowResponseDefault(ctx, fasthttp.StatusUnauthorized, "Company tidak terdaftar")
				} else {
					log.Println(res.Docs[0])
					// models.JsonToStruct(string(res.Docs[0]), &company)
				}
			}
		}
		log.Println(company)
	}
	return ""
}
func GetUserDataToCoreDB(ctx *fasthttp.RequestCtx, idcompany string, username string) models.FindResponse {
	findUserCoreDB := `{"selector":{"$or":[{"_id":"` + idcompany + `"},{"users":"` + username + `"}]}}`

	res, err, code := services.FindDocument(config.GetCredCDBAdmin(), findUserCoreDB, config.DB_CORE_NAME)
	// Jika username terdaftar di DB center / SCM_CORE, login ke DB Company
	// Jika tidak beri notif username tidak terdaftar
	if err != "" {
		models.ShowResponseDefault(ctx, code, err)
	} else {
		if len(res.Docs) == 0 {
			services.ShowResponseDefault(ctx, fasthttp.StatusUnauthorized, "User tidak ditemukan pada server")
		} else {
			var adminDB models.Company
			jsonString := models.StructToJson(res.Docs[0])
			models.JsonToStruct(jsonString, &adminDB)

			// log.Println(adminDB)
			findUserCoreDB = `{"selector":{"table":"user","_id":"` + username + `"},"use_index":"_design/companytable"}`
			resBody, err, code := services.FindDocumentAsComp(adminDB, findUserCoreDB)

			if err != "" {
				services.ShowResponseDefault(ctx, code, err)
			} else {
				var userData models.UserInsert
				if len(resBody.Docs) > 0 {
					jsonStr := models.StructToJson(resBody.Docs[0])
					models.JsonToStruct(jsonStr, &userData)
					log.Println(userData)
				}
			}
		}
		return res
	}
	return res
}
func GenerateJWT(json string, expiredtime int64) string {
	mySigningKey := []byte(config.TOKEN_SALT)
	type Claims struct {
		Json string `json:"data"`
		jwt.StandardClaims
	}
	claims := Claims{json, jwt.StandardClaims{ExpiresAt: expiredtime}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Println(services.StructToJson(models.DefaultResponse{Status: fasthttp.StatusBadRequest, Messege: "GenerateJWT : " + err.Error()}))
	}
	log.Println(ss + "\n")
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
func CheckSession(ctx *fasthttp.RequestCtx) (models.AdminDB, string, string) {
	authHeader := ctx.Request.Header.Peek("Authorization")
	tokenString, err := extractBearerToken(authHeader)
	if err != nil {
		services.ShowResponseDefault(ctx, fasthttp.StatusUnauthorized, err.Error())
		return models.AdminDB{}, "", err.Error()
	} else {
		token, errToken := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.TOKEN_SALT), nil
		})

		if errToken != nil {
			println("\nError token")
			services.ShowResponseDefault(ctx, fasthttp.StatusUnauthorized, errToken.Error())
			return models.AdminDB{}, "", err.Error()
		} else {
			if token.Valid {
				ctx.Response.SetStatusCode(fasthttp.StatusOK)

				var sessionToken models.SessionToken
				claims := token.Claims.(jwt.MapClaims)
				data := claims["data"].(string)
				services.JsonToStruct(string(data), &sessionToken)
				if sessionToken.AdminKey != "" && sessionToken.AdminKey == os.Getenv("API_KEY_ADMIN") {
					urlDB := config.GetCredCDBAdmin()
					print("--You're SuperAdmin--\n" + urlDB)

					return models.AdminDB{}, config.GetCredCDBAdmin(), ""
				} else {
					resRedis, errRedis := services.GetValueRedis(sessionToken.KeyRedis)

					if errRedis != "" {
						models.ShowResponseDefault(ctx, fasthttp.StatusServiceUnavailable, "Error when getting user session, please contact administration")
						return models.AdminDB{}, config.GetCredCDBCompany(os.Getenv("COUCHDB_USER"), os.Getenv("COUCHDB_PASSWORD")), "Error when getting user session, please contact administration"
					}
					if resRedis == "" {
						saveCompanyCredToRedis(sessionToken.IdCompany)
						// models.ShowResponseDefault(ctx, fasthttp.StatusUnauthorized, "User session is not found, please re-login")
						print("\n")
						return models.AdminDB{}, "", "User session is not found, please re-login"
					} else {
						print("--You're Company Admin--\n")

						var sessionFull models.SessionFull
						services.JsonToStruct(resRedis, &sessionFull)
						// log.Println(resRedis)
						// log.Print("DB URL : ", config.GetCredCDBCompany(sessionFull.AdminDB.UserCDB, sessionFull.AdminDB.PassCDB))
						return sessionFull.AdminDB, config.GetCredCDBCompany(sessionFull.AdminDB.UserCDB, sessionFull.AdminDB.PassCDB), ""
					}
				}

			} else {
				services.ShowResponseDefault(ctx, fasthttp.StatusUnauthorized, "Token is invalid")
				return models.AdminDB{}, "", "Token is invalid"
			}
		}
	}
}
func saveCompanyCredToRedis(idcompany string) {
	findCredCompanyDB := `{"selector":{"_id":"` + idcompany + `"}}`
	res, err, code := services.FindDocument(config.GetCredCDBAdmin(), findCredCompanyDB, config.DB_CRED_NAME)
	if err == "" {
		if len(res.Docs) > 0 {

		}
		print(res.Docs, code)
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
