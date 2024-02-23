package validations

import "regexp"

func IsValidUsername(username string) bool {
	usernameRegex := regexp.MustCompile(`[^\w]+$`)
	return usernameRegex.MatchString(username)
}
