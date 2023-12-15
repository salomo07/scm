package main

import (
	"fmt"
	"log"
	"scm/models"
	"scm/routers"
	"scm/services"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

var port = "8080"

func main() {
	// print(config.HashingBcrypt("http://admin:123@10.180.70.75:5984/"))
	// expTime := time.Now().Local().Add(time.Hour*24*30).UnixNano() / 1000

	// go controllers.GenerateJWT([]byte(services.StructToJson(models.AdminCred{AppId: "scm", IdCompany: "c_1702549958278240", UserCDB: "WVhCcGEyVjVMWFl5TFRNeWQyNDBOelpwZFRRelp6aHNkbXRuYlhBM2QzZGpjM016YTJkM2RERTRPREkxWlRRMGJYWTFjelYy", PassCDB: "TjJKbU9UazJObVJsWXpZMVlqVmlOMkUxTVRJM1pUQTJOVFUxWkdRNU5UUT0="})), expTime)

	// go controllers.GenerateJWT([]byte(services.StructToJson(models.AdminCred{AppId: "scm", AdminKey: "$2a$10$vwNlnoWznZXUoW6r6zDGSOwB6H/.Z9WbUC51JYdVJ93BXsF50dHCG"})), expTime)

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
