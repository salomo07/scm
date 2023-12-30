package main

import (
	"fmt"
	"log"
	"scm/config"
	"scm/controllers"
	"scm/models"
	"scm/routers"
	"scm/services"
	"time"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

var port = "8080"

func main() {
	// z, _ := config.EncryptAES("Salomo07")
	// print("\n " + z)
	// hasil, _ := config.DecryptAES("MBJO1avtjKjpc/L6i4FtodEN+FubqVSQnGdqLVvXJS3Fq8m9")
	// print(hasil)
	expTime := time.Now().Local().Add(time.Hour*24*30).UnixNano() / 1000

	// go controllers.GenerateJWT([]byte(services.StructToJson(models.SessionToken{IdAppCompanyUser: "scm*c_1702276535981680*u_1702276535981680", AppId: "scm", IdCompany: "c_1702276535981680"})), expTime)
	user, _ := config.EncryptAES("apikey-v2-32wn476iu43g8lvkgmp7wwcss3kgwt18825e44mv5s5v")
	pass, _ := config.EncryptAES("7bf9966dec65b5b7a5127e06555dd954")
	print(user + "\n")
	print(pass + "\n")
	dec, _ := config.DecryptAES()
	print("\nDec :" + dec)
	go controllers.GenerateJWT(services.StructToJson(models.SessionToken{KeyRedis: "scm*c_1702276535981680*u_34345345", AppId: "scm", IdCompany: "c_1702276535981680"}), expTime)
	// services.SubscribeRedis("c_1702276535981680")
	router := fasthttprouter.New()
	router.GET("/", func(ctx *fasthttp.RequestCtx) {
		services.ShowResponseDefault(ctx, 200, "Welcome to SCM API")
	})
	router.GET("/company/", func(ctx *fasthttp.RequestCtx) {
		services.ShowResponseDefault(ctx, 200, "Ini router /company/")
	})

	routers.Access_MenuRouters(router)
	routers.CompanyRouters(router)
	routers.CouchDBRouters(router)
	routers.UserRouters(router)
	routers.AuthRouters(router)

	router.HandleMethodNotAllowed = false
	router.MethodNotAllowed = func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		ctx.Response.Header.Set("Content-Type", "application/json")
		fmt.Fprintf(ctx, models.StructToJson(models.DefaultResponse{Messege: "Your method is not allowed", Status: fasthttp.StatusMethodNotAllowed}))
	}
	router.NotFound = func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		ctx.Response.Header.Set("Content-Type", "application/json")
		fmt.Fprintf(ctx, models.StructToJson(models.DefaultResponse{Messege: "API is not found", Status: fasthttp.StatusNotFound}))
	}

	server := &fasthttp.Server{Handler: router.Handler}
	go func() {
		if err := server.ListenAndServe(":" + port); err != nil {
			log.Printf("Error starting server: %s\n", err)
		}
	}()
	log.Println("Server listen on :" + port)
	select {}
}
