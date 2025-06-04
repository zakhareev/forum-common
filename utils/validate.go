package utils

import (
	"unicode/utf8"
)

func IsValidUsername(username string) bool {
	return utf8.RuneCountInString(username) >= 3
}
