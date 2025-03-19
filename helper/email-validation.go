package helper

import "regexp"

func IsEmailValid(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.com$`)
	return re.MatchString(email)
}

func IsPasswordValid(password string) bool {
	re := regexp.MustCompile(`[~!@#$%^&*()_+]+`)
	return re.MatchString(password)
}
