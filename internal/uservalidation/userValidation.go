package uservalidation

import "regexp"

const regex = "^[a-zA-Z0-9@._-]*$"

func ValidateInput(input string) bool {
	match, _ := regexp.MatchString(regex, input)
	return match
}
