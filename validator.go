package main

import "regexp"

func validateEmail(email string) bool {
	pattern := "^[a-zA-Z0-9]+@[a-zA-Z0-9]+\\.[a-zA-Z]{2,}$"
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}
