package login

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/natewong1313/UMW-Class-Monitor/http_client"
	"github.com/natewong1313/UMW-Class-Monitor/logger"
)


func (loginTask *LoginTask) submitSSOData(samlResponse, relayState string){
	for{
		logger.Debug("Submitting SSO data")

		resp, err := loginTask.HttpClient.Post(&http_client.Request{
			URL: "https://ssb-prod.ec.umw.edu/ssomanager/saml/SSO",
			Headers: map[string]string{
				"Host": "ssb-prod.ec.umw.edu",
				"Connection": "keep-alive",
				"Cache-Control": "max-age=0",
				"Upgrade-Insecure-Requests": "1",
				"Origin": "https://eis-prod.ec.umw.edu",
				"Content-Type": "application/x-www-form-urlencoded",
				"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36",
				"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
				"Sec-GPC": "1",
				"Sec-Fetch-Site": "same-site",
				"Sec-Fetch-Mode": "navigate",
				"Sec-Fetch-Dest": "document",
				"Referer": "https://eis-prod.ec.umw.edu/",
				"Accept-Encoding": "gzip, deflate, br",
				"Accept-Language": "en-US,en;q=0.9",
			},
			FormBody: url.Values{
				"SAMLResponse": {samlResponse},
				"RelayState": {relayState},
			},
		})
		if err != nil || resp.StatusCode != 200 {
			if resp.StatusCode != 200{
				err = errors.New(fmt.Sprintf("%d response code", resp.StatusCode))
			}
			loginTask.handleErr("Failed to submit SSO data", err)
		}
	
		logger.Debug("Successfully submitted SSO data")
		return 
	}
}
