package header

import "unicode"

const (
	symbol = "！……（），。？、"
)

func isChinese(c int32) bool {
	if unicode.Is(unicode.Han, c) {
		return true
	}

	for _, s := range symbol {
		if c == s {
			return true
		}
	}
	return false
}
