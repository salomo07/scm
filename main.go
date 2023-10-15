package main

import (
	"fmt"
	"log"
	"scm/models"
	"scm/routers"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

var port = "1234"

func main() {
	router := fasthttprouter.New()
	router.GET("/", func(ctx *fasthttp.RequestCtx) {
		fmt.Fprintf(ctx, "Welcome to SCM API")
	})
	routers.CompanyRouters(router)
	routers.CouchDBRouters(router)
	routers.UserRouters(router)
	// print("\nUser: " + config.HashingBcrypt("admin") + "\n")
	// print("\nPass: " + config.HashingBcrypt("123") + "\n")
	// print("\nHost: " + config.HashingBcrypt("10.180.8.74") + "\n")
	// print("\nPort: " + config.HashingBcrypt("5984") + "\n")

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
