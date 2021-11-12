package monitor

import (
	"sync"

	"github.com/natewong1313/UMW-Class-Monitor/http_client"
	"github.com/natewong1313/UMW-Class-Monitor/user_data"
)


func StartMonitorThreads(userData *user_data.UserData, httpClient *http_client.HttpClient){
	var wg sync.WaitGroup

	initialSearch(httpClient)

	monitorStatusWg.Lock()
	
	for i, class := range userData.Classes{
		monitorStatus = append(monitorStatus, []string{class.Subject, class.ClassNumber, class.CRN, "Starting Monitor"})
		go StartMonitor(class, userData.Twilio, httpClient, i)
		wg.Add(1)
	}
	monitorStatusWg.Unlock()

	go StartMonitorStatusDisplay()

	wg.Wait()
}

func StartMonitor(class user_data.Class, twilioData user_data.TwilioData, httpClient *http_client.HttpClient, i int){
	monitorTask := &MonitorTask{
		HttpClient: httpClient,
		TwilioData: twilioData,
		MSIndex: i,
		Subject: class.Subject,
		ClassNumber: class.ClassNumber,
		CRN: class.CRN,
		MonitorDelay: 4000,
		ErrDelay: 4000,
	}
	monitorTask.doSearch()
}