package login

import (
	"github.com/natewong1313/UMW-Class-Monitor/http_client"
	"github.com/natewong1313/UMW-Class-Monitor/logger"
)

func Login(username, password string, httpClient *http_client.HttpClient) bool{
	loginTask := &LoginTask{
		HttpClient: httpClient,
		ErrDelay: 4000,
	}

	samlRequest, relayState := loginTask.fetchSAMLData()
	sessionDataKey := loginTask.submitSAMLData(samlRequest, relayState)
	loginSuccess, samlResponse := loginTask.submitLoginInfo(username, password, sessionDataKey)
	if !loginSuccess{
		return false
	}
	samlResponse, relayState = loginTask.submitCommonAuthData(samlResponse, relayState)
	loginTask.submitSSOData(samlResponse, relayState)
	
	logger.Info("Successfully logged in")
	return true
}