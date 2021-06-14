package util

import (
	"strings"
)

func Capitalize(s string) string {
	if len(s) < 1 {
		return s
	}
	firstLetter := s[0]
	if firstLetter < 65 || (firstLetter > 90 && firstLetter < 97) ||
		firstLetter > 122 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}
