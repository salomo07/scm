package services

import (
	"encoding/json"
	"fmt"
	"scm/models"
	"strconv"
	"time"

	"github.com/valyala/fasthttp"
)

// Generally functional ---->
func JsonToStruct(jsonStr string, dynamic any) interface{} {
	json.Unmarshal([]byte(jsonStr), &dynamic)
	return dynamic
}
func StructToJson(v any) string {
	res, err := json.Marshal(v)
	if err != nil {
		println("Fail to convert to JSON")
	}
	print(string(res))
	return string(res)
}
func ShowResponseDefault(ctx *fasthttp.RequestCtx, statuscode int, msg string) {
	ctx.Response.SetStatusCode(statuscode)
	fmt.Fprintf(ctx, StructToJson(models.DefaultResponse{Status: statuscode, Messege: msg}))
}
func ShowResponseJson(ctx *fasthttp.RequestCtx, statuscode int, jsonString string) {
	ctx.Response.SetStatusCode(statuscode)
	fmt.Fprintf(ctx, jsonString)
}

//<----

func SendToNextServer(url string, method string, body string) (resBody string, errStr string, statuscode int) {
	client := &fasthttp.Client{
		MaxIdleConnDuration: 5 * time.Second,
	}
	forwardedRequest := fasthttp.AcquireRequest()
	// print("\n\nReq to : " + method + " --> " + url)
	forwardedRequest.SetRequestURI(url)
	forwardedRequest.SetBody([]byte(body))
	forwardedRequest.Header.SetMethod(string(method))
	forwardedRequest.Header.SetContentType("application/json")

	forwardedResponse := fasthttp.AcquireResponse()
	err := client.Do(forwardedRequest, forwardedResponse)
	if err != nil {
		print(err.Error())
		return "", err.Error(), fasthttp.StatusInternalServerError
	}
	print("\n" + strconv.Itoa(forwardedResponse.StatusCode()) + " - " + string(forwardedResponse.Body()))
	fasthttp.ReleaseRequest(forwardedRequest)
	// fasthttp.ReleaseResponse(forwardedResponse)
	return string(forwardedResponse.Body()), "", forwardedResponse.StatusCode()
}

func ToCDBCompany(url string, method string, body []byte) (resBody string, errStr string, statuscode int) {
	client := &fasthttp.Client{
		MaxIdleConnDuration: 5 * time.Second,
	}
	forwardedRequest := fasthttp.AcquireRequest()
	forwardedRequest.SetRequestURI(url)
	forwardedRequest.SetBody(body)
	forwardedRequest.Header.SetMethod(string(method))
	forwardedRequest.Header.SetContentType("application/json")

	forwardedResponse := fasthttp.AcquireResponse()
	err := client.Do(forwardedRequest, forwardedResponse)
	if err != nil {
		print(err.Error())
		return "", err.Error(), fasthttp.StatusInternalServerError
	}
	print("\n" + strconv.Itoa(forwardedResponse.StatusCode()) + " - " + string(forwardedResponse.Body()))
	fasthttp.ReleaseRequest(forwardedRequest)
	// fasthttp.ReleaseResponse(forwardedResponse)
	return string(forwardedResponse.Body()), "", fasthttp.StatusOK
}
