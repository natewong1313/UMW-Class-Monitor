package main

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/natewong1313/UMW-Class-Monitor/http_client"
	"github.com/natewong1313/UMW-Class-Monitor/logger"
	"github.com/natewong1313/UMW-Class-Monitor/login"
	"github.com/natewong1313/UMW-Class-Monitor/monitor"
	"github.com/natewong1313/UMW-Class-Monitor/prompts"
	"github.com/natewong1313/UMW-Class-Monitor/user_data"
	"github.com/olekukonko/tablewriter"
)

func main(){
	httpClient, err := http_client.New()
	if err != nil{
		logger.Err("Error creating http client", err)
		return
	}
	
	logger.Info("Initializing UMW Class Monitor")

	var loginSuccess bool
	var hasSetTwilioAccount bool

	userData := user_data.ReadConfigData()
	if userData.Username != "" && userData.Password != ""{
		loginSuccess = login.Login(userData.Username, userData.Password, httpClient)
	}
	if userData.Twilio.AccountSid != "" && userData.Twilio.AuthToken != "" && userData.Twilio.TwilioPhoneNumber != "" && userData.Twilio.PhoneNumber != ""{
		hasSetTwilioAccount = true
	}

	var hasShownBaseInfo bool
	for{
		if !hasShownBaseInfo{
			logger.Debug("----------------------------")
			logger.Debug(fmt.Sprintf("Logged in: %t", loginSuccess))
			logger.Debug(fmt.Sprintf("Classes in monitor list: %d", len(userData.Classes)))
		}
		logger.Debug("----------------------------")
		logger.Info("What action would you like to do today?")
		
		var items []string
		if loginSuccess && hasSetTwilioAccount{
			hasShownBaseInfo = true
			items = []string{"Run monitor", "View classes in monitor list", "Add new class to monitor list", "Remove class from monitor list", "Quit"}
		}else if !loginSuccess && !hasSetTwilioAccount{
			items = []string{"Log in to UMW account", "Set Twilio Account", "Quit"}
		} else if loginSuccess && !hasSetTwilioAccount{
			items = []string{"Set Twilio Account", "Quit"}
		} else if !loginSuccess && hasSetTwilioAccount{
			items = []string{"Log in to UMW account", "Quit"}
		}

		prompt := promptui.Select{
			Label: "Select action",
			Items: items,
		}

		_, result, err := prompt.Run()

		if err != nil {
			if err.Error() != "^C"{
				logger.Err("Error creating prompt", err)
			}
			return
		}

		switch result{
		case "Log in to UMW account":
			username, password, err := prompts.PromptLoginInfo()
			if err == nil{
				loginSuccess = login.Login(username, password, httpClient)
				if loginSuccess{
					userData.Username = username
					userData.Password = password
					user_data.UpdateConfigData(userData)
				}
			}
		case "Set Twilio Account":
			accountSid, authToken, twilioPhoneNumber, phoneNumber, err := prompts.PromptTwilioInfo()
			if err == nil && accountSid != "" && authToken != "" && phoneNumber != ""{
				userData.Twilio.AccountSid = accountSid
				userData.Twilio.AuthToken = authToken
				userData.Twilio.TwilioPhoneNumber = twilioPhoneNumber
				userData.Twilio.PhoneNumber = phoneNumber
				user_data.UpdateConfigData(userData)

				hasSetTwilioAccount = true
			}
		case "Run monitor":
			monitor.StartMonitorThreads(userData, httpClient)
		case "View classes in monitor list":
			logger.Info(fmt.Sprintf("Showing %d classes", len(userData.Classes)))

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Subject", "Class Number", "CRN"})
			for _, class := range userData.Classes{
				table.Append([]string{class.Subject, class.ClassNumber, class.CRN})
			}
			table.Render()
		case "Add new class to monitor list":
			class, err := prompts.PromptAddNewClass()
			if err == nil{
				if user_data.CheckClassInConfig(class, userData){
					logger.Error("Class already in list")
				}else{
					user_data.AddNewClassToConfig(class, userData)
				}
			}
		case "Remove class from monitor list":
			CRN, err := prompts.PromptRemoveClass()
			if err == nil{
				user_data.RemoveClassFromConfig(CRN, userData)
			}
		case "Quit":
			return
		}
	}

}