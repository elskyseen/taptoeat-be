package validations

func IsValidChar(context string, min int, max int) bool {
	return len(context) > min && len(context) < max
}
