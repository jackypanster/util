package util

import (
	"github.com/parnurzeal/gorequest"
	"time"
	"net/http"
	"errors"
	"fmt"
	"bytes"
)

func Post(targetUrl string, content string) (string, error) {
	request := gorequest.New()
	resp, body, errs := request.Post(targetUrl).
		Send(content).
		Retry(3, 7*time.Second, http.StatusBadRequest, http.StatusInternalServerError).End()

	return setupResp(content, resp, body, errs)
}

func setupResp(request string, response *http.Response, body string, errs []error) (string, error) {
	if errs != nil {
		var buffer bytes.Buffer
		buffer.WriteString(fmt.Sprintf("fail to send request %s\n", request))
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
