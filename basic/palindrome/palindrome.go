package palindrome

import "strconv"

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	for i := range str {
		if areOposingCharactersDifferent(i, str) {
			return false
		}
	}
	return true

}

func areOposingCharactersDifferent(i int, str string) bool {
	return str[i] != str[len(str)-1-i]
}
