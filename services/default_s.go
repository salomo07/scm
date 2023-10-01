package services

import (
	"encoding/json"
	"fmt"
	"scm/models"

	"github.com/valyala/fasthttp"
)

func SendToNextServer(ctx *fasthttp.RequestCtx) {
	client := &fasthttp.Client{}
	print(string(ctx.URI().FullURI()))
	forwardedRequest := fasthttp.AcquireRequest()
	forwardedRequest.SetRequestURI(string(ctx.URI().FullURI()))
	forwardedRequest.SetBody(ctx.Request.Body())
	forwardedRequest.Header.SetMethod(string(ctx.Method()))

	forwardedResponse := fasthttp.AcquireResponse()
	err := client.Do(forwardedRequest, forwardedResponse)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadGateway)
		fmt.Fprintf(ctx, StructToJson(models.DefaultResponse{Status: fasthttp.StatusBadRequest, Messege: "idcompany is needed"}))
	}

	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Response.SetStatusCode(forwardedResponse.StatusCode())
	ctx.Response.SetBody(forwardedResponse.Body())
	fasthttp.ReleaseRequest(forwardedRequest)
	fasthttp.ReleaseResponse(forwardedResponse)
}
func JsonToStruct(jsonStr string, dynamic any) interface{} {
	json.Unmarshal([]byte(jsonStr), &dynamic)
	return dynamic
}
func StructToJson(v any) string {
	res, err := json.Marshal(v)
	if err != nil {
		println("Fail to convert to JSON")
	}
	return string(res)
}
