package validations

import "regexp"

func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	return emailRegex.MatchString(email)
}
