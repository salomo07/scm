package main

import (
	"fmt"
	"log"
	"scm/controllers"
	"scm/models"
	"scm/routers"
	"scm/services"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

var port = "8080"

func main() {
	// @Version 1.0.0
	// @Title SCM API
	// @Description API usually works as expected. But sometimes its not true.
	// @ContactName Salomo Sitompul
	// @ContactEmail sitompulsalomo@gmail.com
	// @LicenseName MIT
	// @LicenseURL https://en.wikipedia.org/wiki/MIT_License
	// @Server http://www.fake.com Server-1
	// @Server http://www.fake2.com Server-2
	// @Security AuthorizationHeader read write
	// @SecurityScheme AuthorizationHeader http bearer Input your token
	router := fasthttprouter.New()
	router.GET("/", func(ctx *fasthttp.RequestCtx) {
		services.ShowResponseDefault(ctx, 200, "Welcome to SCM API")
	})
	router.GET("/company/", func(ctx *fasthttp.RequestCtx) {
		services.ShowResponseDefault(ctx, 200, "Ini router /company/")
	})

	// @Title Menambahkan menu aplikasi
	// @Description Get users related to a specific group.
	// @Header models.Menu1
	// @Param  body
	// @Success  200  object  services.ShowResponseDefault(ctx, fasthttp.StatusBadRequest, "Request body cant be empty")  "UsersResponse JSON"
	// @Failure  400  object  ErrorResponse  "ErrorResponse JSON"
	// @Resource users
	// @Route /api/group/{groupID}/users [get]
	router.POST("api/v1/admin/menu/add", func(ctx *fasthttp.RequestCtx) {
		if controllers.CheckSession(ctx) != "" {
			controllers.AddMenu(ctx)
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
	})

	routers.Access_MenuRouters(router)
	routers.CompanyRouters(router)
	routers.CouchDBRouters(router)
	routers.UserRouters(router)

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
