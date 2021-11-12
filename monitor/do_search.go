package monitor

import (
	"fmt"
	"time"

	"github.com/natewong1313/UMW-Class-Monitor/http_client"
	"github.com/tidwall/gjson"
)

func (monitorTask *MonitorTask) doSearch(){
	subjectCodeFound, subjectCode := monitorTask.findSubjectCode()
	if !subjectCodeFound{
		monitorTask.updateStatus("Subject not found")
		return
	}

	lastStatus := 0
	for{
		resp, err := monitorTask.HttpClient.Get(&http_client.Request{
			URL: fmt.Sprintf("https://reg-prod.ec.umw.edu/StudentRegistrationSsb/ssb/searchResults/searchResults?txt_subject=%s&txt_courseNumber=%s&txt_term=202201&startDatepicker=&endDatepicker=&pageOffset=0&pageMaxSize=10&sortColumn=subjectDescription&sortDirection=asc",
							subjectCode, monitorTask.ClassNumber ),
			Headers: map[string]string{
				"accept": "application/json, text/javascript, */*; q=0.01",
				"x-requested-with": "XMLHttpRequest",
				"user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36",
				"sec-gpc": "1",
				"sec-fetch-site": "same-origin",
				"sec-fetch-mode": "cors",
				"sec-fetch-dest": "empty",
				"referer": "https://reg-prod.ec.umw.edu/StudentRegistrationSsb/ssb/classSearch/classSearch",
				"accept-encoding": "gzip, deflate, br",
				"accept-language": "en-US,en;q=0.9",
			},
		})
		if err != nil || resp.StatusCode != 200 {
			time.Sleep(time.Millisecond * time.Duration(monitorTask.ErrDelay))
		}
	
		classFound, classAvailable, seatsAvailable := monitorTask.parseClassStatus(resp.Body)
		
		if !classFound{
			monitorTask.updateStatus("Class Not Found")
		} else if classFound && !classAvailable{
			monitorTask.updateStatus("Class Full")
			lastStatus = 1
		} else if classFound && classAvailable{
			monitorTask.updateStatus(fmt.Sprintf("Class Available (%d Seats)", seatsAvailable))
			if lastStatus == 1 {
				monitorTask.notify(fmt.Sprintf("A class has become available! Subject: %s, course number: %s, CRN: %s", 
									monitorTask.Subject, monitorTask.ClassNumber, monitorTask.CRN))
			}
			lastStatus = 2
		}

		time.Sleep(time.Millisecond * time.Duration(monitorTask.MonitorDelay))
	}
}

func (monitorTask *MonitorTask) parseClassStatus(respBody []byte) (bool, bool, int64){
	for _, classJson := range gjson.Get(string(respBody), "data").Array() {
		if gjson.Get(classJson.String(), "courseReferenceNumber").String() == monitorTask.CRN {
			seatsAvailable := gjson.Get(classJson.String(), "seatsAvailable").Int()
			if seatsAvailable > 0 {
				return true, true, seatsAvailable
			}else{
				return true, false, 0
			}
		}
	}

	return false, false, 0
}