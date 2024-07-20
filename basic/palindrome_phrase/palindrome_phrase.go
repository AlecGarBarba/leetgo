package palindromephrase

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = removeSpecialCharacters(s)
	start, finish := 0, len(s)-1

	for {
		if start >= finish {
			break
		}

		if s[start] != s[finish] {
			return false
		}
		start++
		finish--

	}

	return true

}

func removeSpecialCharacters(s string) string {
	var builder strings.Builder

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			builder.WriteRune(r)
		}
	}

	return strings.ToLower(builder.String())
}
