package prompts

import (
	"errors"
	"strconv"
)

func validateStr(input string) error{
	if len(input) == 0 {
		return errors.New("Input cannot be blank")
	}
	return nil
}

func validateAsInt(input string) error{
	_, err := strconv.Atoi(input)
	if err != nil {
		return errors.New("Input must be a number")
	}
	return nil
}