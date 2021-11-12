package prompts

import (
	"github.com/manifoldco/promptui"
	"github.com/natewong1313/UMW-Class-Monitor/logger"
)

func PromptLoginInfo() (string, string, error){
	prompt := promptui.Prompt{
		Label:    "Username",
		Validate: validateStr,
	}
	username, err := prompt.Run()
	if err != nil {
		if err.Error() != "^C"{
			logger.Err("Error creating prompt", err)
		}
		return "", "", err
	}

	prompt = promptui.Prompt{
		Label:    "Password",
		Validate: validateStr,
		Mask:     '*',
	}
	password, err := prompt.Run()
	if err != nil {
		if err.Error() != "^C"{
			logger.Err("Error creating prompt", err)
		}
		return "", "", err
	}

	return username, password, nil
}