package services

import (
	"encoding/json"
	"fmt"
	"scm/models"
	"strconv"

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
	return string(res)
}
func ShowResponseDefault(ctx *fasthttp.RequestCtx, statuscode int, msg string) {
	fmt.Fprintf(ctx, StructToJson(models.DefaultResponse{Status: statuscode, Messege: msg}))
}

//<----

func SendToNextServer(url string, method string, body []byte) (resBody string, errStr string) {
	client := &fasthttp.Client{}
	forwardedRequest := fasthttp.AcquireRequest()
	forwardedRequest.SetRequestURI(url)
	forwardedRequest.SetBody(body)
	forwardedRequest.Header.SetMethod(string(method))
	forwardedRequest.Header.SetContentType("application/json")

	forwardedResponse := fasthttp.AcquireResponse()
	err := client.Do(forwardedRequest, forwardedResponse)
	if err != nil {
		print(err.Error())
		return "", err.Error()
	}
	print("\n" + strconv.Itoa(forwardedResponse.StatusCode()) + " - " + string(forwardedResponse.Body()))
	// ctx.Response.Header.Set("Content-Type", "application/json")
	// ctx.Response.SetStatusCode(forwardedResponse.StatusCode())
	// ctx.Response.SetBody(forwardedResponse.Body())
	fasthttp.ReleaseRequest(forwardedRequest)
	// fasthttp.ReleaseResponse(forwardedResponse)
	return string(forwardedResponse.Body()), ""

}