package util

import (
	"regexp"

)

// IsContainSpecialCharacters Validade '[^a-zA-Z0-9 ]+' regex.
func IsContainSpecialCharacters(value string) (bool, error) {
	cleaner, err := ReplaceCharacters(value, "[^a-zA-Z0-9 ]+")
	if err != nil {
		return false, err
	}

	if len(cleaner) != len(value) {
		return true, nil
	}

	return false, nil
}

// IsContainOnlyNumbers Validade '[^0-9]+' regex.
func IsContainOnlyNumbers(value string) (bool, error) {
	cleaner, err := ReplaceCharacters(value, "[^0-9]+")
	if err != nil {
		return false, err
	}

	if len(cleaner) != len(value) {
		return true, nil
	}

	return false, nil
}

// IsContainOnlyLetters Validade '[^0-9 ]+' regex.
func IsContainOnlyLetters(value string) (bool, error) {
	cleaner, err := ReplaceCharacters(value, "[^a-zA-Z ]+")
	if err != nil {
		return false, err
	}

	if len(cleaner) != len(value) {
		return true, nil
	}

	return false, nil
}

// ReplaceCharacters ...
func ReplaceCharacters(value string, regex string) (string, error) {
	if value == "" {
		return value, nil
	}

	reg, err := regexp.Compile(regex)
	if err != nil {
		return "", err
	}
	cleaner := reg.ReplaceAllString(value, "")

	return cleaner, nil
}
