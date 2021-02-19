package helper

import (
	"fmt"
	"strings"

	"gopkg.in/go-playground/validator.v9"

)

// ValidateRequestContract Based on the 'https://dev.to/gayanhewa/request-validation-with-go-59a3' solution.
func ValidateRequestContract(request interface{}) (bool, map[string]string) {
	errors := make(map[string]string)
	v := validator.New()
	v.SetTagName("validate")
	if err := v.Struct(request); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err.Tag() == "email" {
				errors[strings.ToLower(err.Field())] = "Invalid E-mail format."
				continue
			}
			errors[strings.ToLower(err.Field())] = fmt.Sprintf("%s is %s %s", err.Field(), err.Tag(), err.Param())
		}
		return false, errors
	}
	return true, nil
}
