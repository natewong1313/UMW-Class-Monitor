package login

import (
	"bytes"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/natewong1313/UMW-Class-Monitor/http_client"
	"github.com/natewong1313/UMW-Class-Monitor/logger"
)


func (loginTask *LoginTask) submitLoginInfo(username, password, sessionDataKey string) (bool, string) {
	for{
		logger.Debug("Submitting login info")

		resp, err := loginTask.HttpClient.Post(&http_client.Request{
			URL: "https://auth.umw.edu/samlsso",
			Headers: map[string]string{
				"Host": "auth.umw.edu",
				"Connection": "keep-alive",
				"Cache-Control": "max-age=0",
				"Upgrade-Insecure-Requests": "1",
				"Origin": "https://auth.umw.edu",
				"Content-Type": "application/x-www-form-urlencoded",
				"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36",
				"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
				"Sec-GPC": "1",
				"Sec-Fetch-Site": "same-origin",
				"Sec-Fetch-Mode": "navigate",
				"Sec-Fetch-User": "?1",
				"Sec-Fetch-Dest": "document",
				"Accept-Encoding": "gzip, deflate, br",
				"Accept-Language": "en-US,en;q=0.9",
			},
			FormBody: url.Values{
				"tocommonauth": {"true"},
				"username": {username},
				"password": {password},
				"sessionDataKey": {sessionDataKey},
			},
		})
		if err != nil || resp.StatusCode != 200 {
			if resp.StatusCode != 200{
				err = errors.New(fmt.Sprintf("%d response code", resp.StatusCode))
			}
			loginTask.handleErr("Failed to submit login info", err)
		}
	
		if strings.Contains(string(resp.Body), "Authentication failed! Please try again."){
			logger.Error("Invalid login information")
			return false, ""
		}else{
			samlResponse, err := parseSAMLResponseData(resp.Body)
			if err != nil {
				loginTask.handleErr("Failed to parse SAML response data", err)
			}

			logger.Debug("Successfully submitted login info")
			return true, samlResponse
		}
	}
}

func parseSAMLResponseData(respBody []byte) (string, error){
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(respBody))
	if err != nil{
		return "", err
	}

	samlResponse := doc.Find(`[name="SAMLResponse"]`).AttrOr("value", "")
	return samlResponse, nil
}