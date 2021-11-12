package login

import "github.com/natewong1313/UMW-Class-Monitor/http_client"

type LoginTask struct {
	HttpClient *http_client.HttpClient
	ErrDelay int
}