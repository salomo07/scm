package main

import (
	"fmt"
	"log"
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
	expTime := time.Now().Local().Add(time.Hour*24*30).UnixNano() / 1000
	jsonToken := services.StructToJson(models.Session{AppId: "scm", AdminKey: "$2a$10$vwNlnoWznZXUoW6r6zDGSOwB6H/.Z9WbUC51JYdVJ93BXsF50dHCG", IdUser: "u_000000001"})
	go controllers.GenerateJWT(jsonToken, expTime)

	// go controllers.GenerateJWT(services.StructToJson(models.SessionToken{KeyRedis: "scm*c_1704987640641606*7a31b1142ea980c599b29c213e77c196", AppId: "scm", IdCompany: "c_1704987640641606"}), expTime)
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
