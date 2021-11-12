package http_client

import (
	"io"
	"net/http"
	"strings"
)

func (httpClient *HttpClient) Get(reqData *Request) (*Response, error) {
	reqData.Method = "GET"
	return httpClient.Do(reqData)
}

func (httpClient *HttpClient) Post(reqData *Request) (*Response, error) {
	reqData.Method = "POST"
	return httpClient.Do(reqData)
}

func (httpClient *HttpClient) Do(reqData *Request) (*Response, error) {

	var reqBody io.Reader
	if len(reqData.FormBody) > 0 {
		reqBody = strings.NewReader(reqData.FormBody.Encode())
	}

	req, err := http.NewRequest(reqData.Method, reqData.URL, reqBody)
	if err != nil {
		return nil, err
	}

	for key, val := range reqData.Headers {
		req.Header.Set(key, val)
	}

	resp, err := httpClient.BaseClient.Do(req)
	if err != nil {
		return nil, err
	}
	return buildResponse(resp)
}