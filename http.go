package util

import (
	"bytes"
	"errors"
	"net/http"
	"time"

	"github.com/parnurzeal/gorequest"
)

func Send(request *gorequest.SuperAgent, targetUrl string, content string) (string, error) {
	//request := gorequest.New().Timeout(200 * time.Millisecond)
	//request := gorequest.New()
	//request.Transport.DisableKeepAlives = true
	//request.SetDebug(debug)
	resp, body, errs := request.Post(targetUrl).Send(content).
		Retry(1, time.Second, http.StatusGatewayTimeout, http.StatusRequestTimeout, http.StatusInternalServerError).End()
	return setupResp(content, resp, body, errs)
}

func Post(targetUrl string, content string, debug bool) (string, error) {
	//request := gorequest.New().Timeout(200 * time.Millisecond)
	request := gorequest.New()
	request.Transport.DisableKeepAlives = true
	request.SetDebug(debug)
	resp, body, errs := request.Post(targetUrl).Send(content).
		Retry(1, time.Second, http.StatusGatewayTimeout, http.StatusRequestTimeout, http.StatusInternalServerError).End()
	return setupResp(content, resp, body, errs)
}

func Get(targetUrl string, debug bool) (string, error) {
	//request := gorequest.New().Timeout(200 * time.Millisecond)
	request := gorequest.New()
	request.Transport.DisableKeepAlives = true
	request.SetDebug(debug)
	resp, body, errs := request.Get(targetUrl).
		Retry(1, time.Second, http.StatusGatewayTimeout, http.StatusRequestTimeout, http.StatusInternalServerError).End()
	return setupResp(targetUrl, resp, body, errs)
}

func setupResp(request string, response *http.Response, body string, errs []error) (string, error) {
	if errs != nil {
		var buffer bytes.Buffer
		if len(errs) > 0 {
			for _, err := range errs {
				buffer.WriteString(err.Error())
			}
		}
		return "", errors.New(buffer.String())
	} else if response.StatusCode != http.StatusOK {
		return "", errors.New(response.Status)
	}
	return body, nil
}
