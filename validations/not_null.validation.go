package validations

func IsNotNull(context string) bool {
	return context == "" || len(context) == 0
}
