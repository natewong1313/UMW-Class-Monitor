package monitor

import (
	"net/url"
	"time"

	"github.com/natewong1313/UMW-Class-Monitor/http_client"
)

func initialSearch(httpClient *http_client.HttpClient){
	for{
		resp, err := httpClient.Post(&http_client.Request{
			URL: "https://reg-prod.ec.umw.edu/StudentRegistrationSsb/ssb/term/search?mode=search",
			Headers: map[string]string{
				"accept": "*/*",
				"x-requested-with": "XMLHttpRequest",
				"user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36",
				"content-type": "application/x-www-form-urlencoded; charset=UTF-8",
				"sec-gpc": "1",
				"origin": "https://reg-prod.ec.umw.edu",
				"sec-fetch-site": "same-origin",
				"sec-fetch-mode": "cors",
				"sec-fetch-dest": "empty",
				"referer": "https://reg-prod.ec.umw.edu/StudentRegistrationSsb/ssb/term/termSelection?mode=search",
				"accept-encoding": "gzip, deflate, br",
				"accept-language": "en-US,en;q=0.9",
			},
			FormBody: url.Values{
				"term": {"202201"},
				"studyPath": {""},
				"studyPathText": {""},
				"startDatepicker": {""},
				"endDatepicker": {""},
				"uniqueSessionId": {""},
			},
		})
		if err != nil || resp.StatusCode != 200 {
			time.Sleep(time.Millisecond * time.Duration(4000))
		}
	
		return 
	}
}