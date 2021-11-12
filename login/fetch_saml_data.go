package login

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/natewong1313/UMW-Class-Monitor/http_client"
	"github.com/natewong1313/UMW-Class-Monitor/logger"
)

func (loginTask *LoginTask) fetchSAMLData() (string, string){
	for{
		logger.Debug("Fetching SAML data")

		resp, err := loginTask.HttpClient.Get(&http_client.Request{
			URL: "https://ssb-prod.ec.umw.edu/ssomanager/saml/login?relayState=/c/auth/SSB",
			Headers: map[string]string{
				"Host": "ssb-prod.ec.umw.edu",
				"Connection": "keep-alive",
				"Upgrade-Insecure-Requests": "1",
				"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36",
				"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
				"Sec-GPC": "1",
				"Sec-Fetch-Site": "same-site",
				"Sec-Fetch-Mode": "navigate",
				"Sec-Fetch-User": "?1",
				"Sec-Fetch-Dest": "document",
				"Referer": "https://technology.umw.edu/",
				"Accept-Encoding": "gzip, deflate, br",
				"Accept-Language": "en-US,en;q=0.9",
			},
		})
		if err != nil || resp.StatusCode != 200 {
			if resp.StatusCode != 200{
				err = errors.New(fmt.Sprintf("%d response code", resp.StatusCode))
			}
			loginTask.handleErr("Failed to fetch SAML data", err)
		}
	
		samlRequest, relayState, err := parseInitialSAMLData(resp.Body)
		if err != nil {
			loginTask.handleErr("Failed to parse SAML data", err)
		}
		logger.Debug("Successfully fetched SAML data")
		return samlRequest, relayState	
	}
}

func parseInitialSAMLData(respBody []byte) (string, string, error){
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(respBody))
	if err != nil{
		return "", "", err
	}

	samlRequest := doc.Find(`[name="SAMLRequest"]`).AttrOr("value", "")
	relayState := doc.Find(`[name="RelayState"]`).AttrOr("value", "")
	return samlRequest, relayState, nil
}