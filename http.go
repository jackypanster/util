package util

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/parnurzeal/gorequest"
)

func Post(targetUrl string, content string, debug bool) (string, error) {
	request := gorequest.New().Timeout(200 * time.Millisecond)
	request.SetDebug(debug)
	resp, body, errs := request.Post(targetUrl).Send(content).
		Retry(1, 1*time.Second, http.StatusGatewayTimeout, http.StatusRequestTimeout, http.StatusBadRequest, http.StatusInternalServerError).End()

	if errs != nil && len(errs) > 0 {
		log.Printf("error %+v\n", errs[0])
	}
	return setupResp(content, resp, body, errs)
}

func Get(targetUrl string, debug bool) (string, error) {
	request := gorequest.New().Timeout(200 * time.Millisecond)
	request.SetDebug(debug)
	resp, body, errs := request.Get(targetUrl).
		Retry(1, 1*time.Second, http.StatusGatewayTimeout, http.StatusRequestTimeout, http.StatusBadRequest, http.StatusInternalServerError).End()
	if errs != nil && len(errs) > 0 {
		log.Printf("error %+v\n", errs[0])
	}
	return setupResp(targetUrl, resp, body, errs)
}

func setupResp(request string, response *http.Response, body string, errs []error) (string, error) {
	if errs != nil {
		var buffer bytes.Buffer
		buffer.WriteString(fmt.Sprintf("fail to make HTTP request %s\n", request))
		if len(errs) > 0 {
			for err := range errs {
				buffer.WriteString(fmt.Sprintf("%+v\n", err))
			}
		}
		return "", errors.New(buffer.String())
	} else if response.StatusCode != http.StatusOK {
		return "", errors.New(response.Status)
	}
	return body, nil
}
