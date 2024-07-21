package wordpattern

import "strings"

/*
NOTE: we can make it more readable by using runes instead of stirngs
and by refactoring the if statements. But I'm looking for clarity in this example
*/
func wordPattern(pattern string, s string) bool {

	words := strings.Split(s, " ")

	if len(words) != len(pattern) {
		return false
	}

	mPD := make(map[string]string)
	mDP := make(map[string]string)

	for i, r := range pattern {
		valPD, okPD := mPD[string(r)]
		valDP, okDP := mDP[words[i]]

		if okPD && valPD != words[i] {
			return false
		}

		if okDP && valDP != string(r) {
			return false
		}

		mPD[string(r)] = words[i]
		mDP[words[i]] = string(r)
	}

	return true
}
