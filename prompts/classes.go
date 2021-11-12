package prompts

import (
	"github.com/manifoldco/promptui"
	"github.com/natewong1313/UMW-Class-Monitor/logger"
	"github.com/natewong1313/UMW-Class-Monitor/user_data"
)

func PromptAddNewClass() (user_data.Class, error){
	prompt := promptui.Prompt{
		Label:    "Subject",
		Validate: validateStr,
	}
	subject, err := prompt.Run()
	if err != nil {
		if err.Error() != "^C"{
			logger.Err("Error creating prompt", err)
		}
		return user_data.Class{}, err
	}

	prompt = promptui.Prompt{
		Label:    "Class Number",
		Validate: validateAsInt,
	}
	classNumber, err := prompt.Run()
	if err != nil {
		if err.Error() != "^C"{
			logger.Err("Error creating prompt", err)
		}
		return user_data.Class{}, err
	}

	prompt = promptui.Prompt{
		Label:    "CRN",
		Validate: validateAsInt,
	}
	CRN, err := prompt.Run()
	if err != nil {
		if err.Error() != "^C"{
			logger.Err("Error creating prompt", err)
		}
		return user_data.Class{}, err
	}

	return user_data.Class{
		Subject: subject,
		ClassNumber: classNumber,
		CRN: CRN,
	}, nil
}

func PromptRemoveClass()(string, error){
	prompt := promptui.Prompt{
		Label:    "CRN",
		Validate: validateAsInt,
	}
	CRN, err := prompt.Run()
	if err != nil {
		if err.Error() != "^C"{
			logger.Err("Error creating prompt", err)
		}
		return "", err
	}

	return CRN, nil
}