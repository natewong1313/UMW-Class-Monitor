package http_client

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"
)

func buildResponse(resp *http.Response) (*Response, error) {
	respBody, err := parseBody(resp)
	if err != nil {
		return nil, err
	}
	return &Response{
		StatusCode: resp.StatusCode,
		URL:        resp.Request.URL.String(),
		Body:       respBody,
		Cookies:    resp.Cookies(),
		Header:     resp.Header,
	}, nil
}

func parseBody(resp *http.Response) ([]byte, error) {
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ = gzip.NewReader(resp.Body)
		defer reader.Close()
	default:
		reader = resp.Body
	}

	bodyBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	return bodyBytes, nil
}