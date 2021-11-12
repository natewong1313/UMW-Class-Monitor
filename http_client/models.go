package http_client

import (
	"net/http"
	"net/url"
)

type HttpClient struct {
	BaseClient *http.Client 
}

type Request struct {
	URL         string
	Method      string
	Headers     map[string]string
	FormBody    url.Values
}

type Response struct {
	StatusCode int
	URL        string
	Body       []byte
	Cookies    []*http.Cookie
	Header     http.Header
}