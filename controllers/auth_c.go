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
	"time"

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
			print("Company tidak ditemukan di Redis\n")
			queryFindCred := `{"selector":{"_id":"` + loginInput.IdCompany + `","appid":"` + loginInput.AppId + `"}}`
			res, err, code := services.FindDocument(config.GetCredCDBAdmin(), queryFindCred, config.DB_CORE_NAME)
			if err != "" {
				models.ShowResponseDefault(ctx, fasthttp.StatusInternalServerError, err+"Error code : "+strconv.Itoa(code))
			} else {
				if len(res.Docs) == 0 {
					services.ShowResponseDefault(ctx, fasthttp.StatusUnauthorized, "Company tidak terdaftar")
				} else {
					print("Companynya :\n")
					json := models.StructToJson(res.Docs[0])

					log.Println(json)
					models.JsonToStruct(json, &company)
					// Jika di Redis belum ada data company, maka data yang didapat dari DB
					go services.SaveValueRedis(company.IdCompany, models.StructToJson(company))
					FindUserOnCoreDB(ctx, loginInput)
				}
			}
		} else {
			print("Ini Hasil dari Redis\n")
			//Data Company berhasil ditemukan di Redis
			print(companyData)
			FindUserOnCoreDB(ctx, loginInput)
			// GenerateJWT()
		}
		// log.Println(company)
	}
	return ""
}
func FindUserOnCoreDB(ctx *fasthttp.RequestCtx, loginInput models.LoginInput) {
	query := `{"selector":{"$or": [{"username":"` + loginInput.Username + `","idcompany":"` + loginInput.IdCompany + `","table":"user"},{"email":"` + loginInput.Username + `","idcompany":"` + loginInput.IdCompany + `","table":"user"}]}}`
	findRes, err, code := services.FindDocument(config.GetCredCDBAdmin(), query, config.DB_CORE_NAME)
	if err != "" {
		models.ShowResponseDefault(ctx, fasthttp.StatusInternalServerError, "An error occurred (Error : "+strconv.Itoa(code)+" - "+err+")")
	} else {
		if len(findRes.Docs) == 0 {
			services.ShowResponseDefault(ctx, fasthttp.StatusUnauthorized, "Username or Email is not found")
		} else {
			//Decrypt Pass
			jsonUserSuccess := ""
			for _, val := range findRes.Docs {
				var userData models.User
				json := models.StructToJson(val)
				models.JsonToStruct(json, &userData)
				if config.CompareHashAndPasswordBcrypt(loginInput.Password, userData.Password) != "" {
					user := models.RemoveField(val, "password")
					jsonUserSuccess = models.StructToJson(user)
				}
			}
			if jsonUserSuccess == "" {
				models.ShowResponseDefault(ctx, fasthttp.StatusUnauthorized, "Username, Email or Password is not match")
			} else {

				expTime := time.Now().Local().Add(time.Hour*4).UnixNano() / 1000
				var userJWT models.User
				models.JsonToStruct(jsonUserSuccess, &userJWT)

				jwt := GenerateJWT(models.StructToJson(userJWT), expTime)
				services.SaveValueRedis(loginInput.AppId+"*"+userJWT.IdCompany+"*"+userJWT.Username, jwt, strconv.FormatInt(expTime, 10))

				services.ShowResponseJson(ctx, fasthttp.StatusOK, models.StructToJson(&models.LoginResponse{AppId: loginInput.AppId, IdCompany: userJWT.IdCompany, Token: jwt, Expired: time.Unix(0, expTime*int64(time.Microsecond)).Format(time.RFC3339)}))

			}
		}
	}
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
				var userData models.User
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
func getCompanyDataOnRedisOrDB(ctx *fasthttp.RequestCtx, sessionToken models.Session) (company models.Company, urlDB string, err string) {
	resRedis, errRedis := services.GetValueRedis(sessionToken.KeyRedis)
	var companyModel models.Company
	if errRedis != "" {
		models.ShowResponseDefault(ctx, fasthttp.StatusServiceUnavailable, "Error when getting user session, please contact administration")
		return companyModel, "", "Error when getting user session, please contact administration"
	} else if resRedis == "" {
		print("\n")
		//Jika token yng dimasukan untuk akses Company, tapi token tersebut tidak ditemukan di Redis.
		//Find credDB company by Idcompany
		query := `{"selector":{"_id":"` + sessionToken.IdCompany + `","table":"company"}}`
		find, err, _ := services.FindDocument(config.GetCredCDBAdmin(), query, config.DB_CORE_NAME)
		if err == "" {
			if len(find.Docs) > 0 {
				println(find.Docs[0])
				json := models.StructToJson(find.Docs[0])
				models.JsonToStruct(json, &companyModel)
				return companyModel, config.GetCredCDBCompany(companyModel.UserCDB, companyModel.PassCDB), ""
			} else {
				models.ShowResponseDefault(ctx, fasthttp.StatusUnauthorized, "Company is unregistered")
				return companyModel, "", "Company is unregistered"
			}
		} else {
			return companyModel, "", err
		}
	} else {
		println("\nApa bisa lewat sini???\n")
		return companyModel, "", err
	}
}
func CheckSession(ctx *fasthttp.RequestCtx) (admReturn models.AdminDB, urldb string, errString string) {
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

				var sessionModel models.Session
				claims := token.Claims.(jwt.MapClaims)
				data := claims["data"].(string)
				services.JsonToStruct(string(data), &sessionModel)
				if sessionModel.AdminKey != "" && sessionModel.AdminKey == os.Getenv("API_KEY_ADMIN") {
					//Jika token yang diberikan token SuperAdmin
					err := models.ValidateRequiredFields(sessionModel, ctx)
					urlDB := config.GetCredCDBAdmin()
					print("--You're SuperAdmin--\n" + urlDB)
					log.Println(err, sessionModel)
					company, url, err := getCompanyDataOnRedisOrDB(ctx, sessionModel)
					if err != "" {
						return models.AdminDB{UserCDB: company.UserCDB, PassCDB: company.PassCDB}, config.GetCredCDBAdmin(), ""
					}
					print(url)
				} else {
					//Token sebagai Company
					// getCompanyDataOnRedisOrDB(ctx, sessionModel)
				}

			} else {
				services.ShowResponseDefault(ctx, fasthttp.StatusUnauthorized, "Token is invalid")
				return models.AdminDB{}, "", "Token is invalid"
			}
		}
	}
	return models.AdminDB{}, "", "Token is invalid"
}

// func saveCompanyCredToRedis(idcompany string) {
// 	findCredCompanyDB := `{"selector":{"_id":"` + idcompany + `"}}`
// 	res, err, code := services.FindDocument(config.GetCredCDBAdmin(), findCredCompanyDB, config.DB_CRED_NAME)
// 	if err == "" {
// 		if len(res.Docs) > 0 {

//			}
//			print(res.Docs, code)
//		}
//	}
func extractBearerToken(authHeader []byte) (string, error) {
	// Check if the Authorization header starts with "Bearer "
	if !strings.HasPrefix(string(authHeader), "Bearer ") {
		return "", fmt.Errorf("invalid Bearer token format")
	}

	// Extract the token (remove "Bearer " prefix)
	token := strings.TrimPrefix(string(authHeader), "Bearer ")

	return token, nil
}
