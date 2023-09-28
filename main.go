package main

import (
	"fmt"
	"log"
	"scm/controllers"
	"scm/models"
	"scm/routers"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

var port = "1234"

func main() {
	router := fasthttprouter.New()
	routers.UsersRouters(router)

	router.HandleMethodNotAllowed = false
	router.MethodNotAllowed = func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		res := models.DefaultResponse{Messeges: "Your method is not allowed", Status: fasthttp.StatusMethodNotAllowed}
		fmt.Fprintf(ctx, controllers.StructToJson(res))
	}
	router.NotFound = func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		res := models.DefaultResponse{Messeges: "API is not found", Status: fasthttp.StatusNotFound}
		fmt.Fprintf(ctx, controllers.StructToJson(res))
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
