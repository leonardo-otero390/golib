package golib

import (
	"strings"
)

func IsPalindrome(str string) bool {
	withoutSpaces := strings.ReplaceAll(str, " ", "")
	strLen := len(withoutSpaces)
	for i := 0; i < strLen/2; i++ {
		if withoutSpaces[i] != withoutSpaces[strLen-i-1] {
			return false
		}
	}
	return true
}
