package utils

import (
	"regexp"
	"strings"
)

func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func ValidateName(name string) bool {
	return len(strings.TrimSpace(name)) >= 2
}

func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
