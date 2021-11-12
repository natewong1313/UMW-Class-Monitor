package monitor

import (
	"github.com/natewong1313/UMW-Class-Monitor/http_client"
	"github.com/natewong1313/UMW-Class-Monitor/user_data"
)

type MonitorTask struct {
	HttpClient *http_client.HttpClient
	TwilioData user_data.TwilioData
	// Monitor Status Slice Index
	MSIndex int
	Subject string
	ClassNumber string
	CRN string
	MonitorDelay int
	ErrDelay int
}