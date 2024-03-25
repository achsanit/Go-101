package helper

import "regexp"

func IsValidEmail(email string) bool {
	// regex pattern email
    pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

    // Compile the regular expression
    regex := regexp.MustCompile(pattern)

    // Check if the email matches the pattern
    return regex.MatchString(email)
}