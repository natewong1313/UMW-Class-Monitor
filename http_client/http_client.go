package http_client

import (
	"net/http"
	"net/http/cookiejar"
)

func New() (*HttpClient, error){
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	return &HttpClient{
		BaseClient: &http.Client{
			Jar:       jar,
		},
	}, nil
}