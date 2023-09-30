package services

import (
	"scm/config"

	"github.com/valyala/fasthttp"
)

func CreateDB(userdb string, passdb string, dbname string) {
	url := config.GetCredCDB(userdb, passdb) + dbname
	method := "PUT"
	RequestToCDB(url, method, nil)
}

func RequestToCDB(url string, method string, body []byte) {
	client := &fasthttp.Client{}
	forwardedRequest := fasthttp.AcquireRequest()
	forwardedRequest.SetRequestURI(url)
	// forwardedRequest.SetBody(ctx.Request.Body())
	forwardedRequest.SetBody(body)
	forwardedRequest.Header.SetMethod(string(method))

	forwardedResponse := fasthttp.AcquireResponse()
	err := client.Do(forwardedRequest, forwardedResponse)
	if err != nil {
		println("Error sending : " + err.Error())
	}
	// ctx.Response.Header.Set("Content-Type", "application/json")
	// ctx.Response.SetStatusCode(forwardedResponse.StatusCode())
	// ctx.Response.SetBody(forwardedResponse.Body())
	fasthttp.ReleaseRequest(forwardedRequest)
	fasthttp.ReleaseResponse(forwardedResponse)
}
