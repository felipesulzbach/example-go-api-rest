package contract

import (
	"errors"
	"time"

	"github.com/felipesulzbach/example-go-api-rest/src/domain/util"

)

func validateStringOnlyNumbers(param string, value string) error {
	response, _ := util.IsContainOnlyNumbers(value)
	if response {
		return errors.New("The " + param + " parameter must contain only numbers")
	}

	return nil
}

func validateStringSpecialCharacter(param string, value string) error {
	response, _ := util.IsContainSpecialCharacters(value)
	if response {
		return errors.New("The " + param + " parameter must contain only numbers and letters")
	}

	return nil
}

func validateID(id int64) error {
	if id == 0 {
		return errors.New("The ID parameter must be informed and its value must be greater than zero")
	}

	return nil
}

func validateDateIsDateTimeEmpty(param string, value time.Time) error {
	isEmpty := util.IsDateTimeEmpty(value)
	if isEmpty {
		return errors.New("The " + param + " parameter must be informed")
	}

	return nil
}

func validadeStarDateAfterEndDate(valueStart time.Time, valueEnd time.Time) error {
	isBefore := util.IsDateTimeBefore(valueStart, valueEnd)
	if !isBefore {
		return errors.New("The start date cannot be greater than the end date")
	}

	return nil
}
