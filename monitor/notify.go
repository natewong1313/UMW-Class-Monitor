package monitor

import (
	"os"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func (monitorTask *MonitorTask) notify(msg string){
	os.Setenv("TWILIO_ACCOUNT_SID", monitorTask.TwilioData.AccountSid)
	os.Setenv("TWILIO_AUTH_TOKEN", monitorTask.TwilioData.AuthToken)

	client := twilio.NewRestClient()

	params := &openapi.CreateMessageParams{}
    params.SetTo(monitorTask.TwilioData.PhoneNumber)
    params.SetFrom(monitorTask.TwilioData.TwilioPhoneNumber)
    params.SetBody(msg)

    client.ApiV2010.CreateMessage(params)
}