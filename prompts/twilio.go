package prompts

import (
	"github.com/manifoldco/promptui"
	"github.com/natewong1313/UMW-Class-Monitor/logger"
)

func PromptTwilioInfo() (string, string, string, string, error){
	prompt := promptui.Prompt{
		Label:    "Account Sid",
		Validate: validateStr,
	}
	accountSid, err := prompt.Run()
	if err != nil {
		if err.Error() != "^C"{
			logger.Err("Error creating prompt", err)
		}
		return "", "", "", "", err
	}

	prompt = promptui.Prompt{
		Label:    "Auth Token",
		Validate: validateStr,
	}
	authToken, err := prompt.Run()
	if err != nil {
		if err.Error() != "^C"{
			logger.Err("Error creating prompt", err)
		}
		return "", "", "", "", err
	}

	prompt = promptui.Prompt{
		Label:    "Twilio Phone Number",
		Validate: validateAsInt,
	}
	twilioPhoneNumber, err := prompt.Run()
	if err != nil {
		if err.Error() != "^C"{
			logger.Err("Error creating prompt", err)
		}
		return "", "", "", "", err
	}

	prompt = promptui.Prompt{
		Label:    "Your Phone Number",
		Validate: validateAsInt,
	}
	phoneNumber, err := prompt.Run()
	if err != nil {
		if err.Error() != "^C"{
			logger.Err("Error creating prompt", err)
		}
		return "", "", "", "", err
	}

	return accountSid, authToken, twilioPhoneNumber, phoneNumber, nil
}