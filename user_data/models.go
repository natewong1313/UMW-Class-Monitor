package user_data

type UserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Twilio TwilioData `json:"twilio"`
	Classes []Class `json:"classes"`
}

type TwilioData struct {
	AccountSid string `json:"accountSid"`
	AuthToken string `json:"authToken"`
	TwilioPhoneNumber string `json:"twilioPhoneNumber"`
	PhoneNumber string `json:"phoneNumber"`
}

type Class struct {
	Subject string `json:"subject"`
	ClassNumber string `json:"classNumber"`
	CRN string `json:"CRN"`
}