package roman

import "strings"

/**
Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given a roman numeral, convert it to an integer.

Constraints:

1 <= s.length <= 15
s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
It is guaranteed that s is a valid roman numeral in the range [1, 3999].
*/

func romanToInt(s string) int {
	rm := romanMap()
	s = replaceEdgeCases(s)
	normal := 0
	for _, r := range s {
		c := string(r)
		normal += rm[c]
	}

	return normal
}

func replaceEdgeCases(s string) string {
	s = strings.Replace(s, "IV", "IIII", -1)
	s = strings.Replace(s, "IX", "VIIII", -1)
	s = strings.Replace(s, "XL", "XXXX", -1)
	s = strings.Replace(s, "XC", "LXXXX", -1)
	s = strings.Replace(s, "CD", "CCCC", -1)
	return strings.Replace(s, "CM", "DCCCC", -1)
}

func romanMap() map[string]int {
	romansMap := make(map[string]int)
	romansMap["I"] = 1
	romansMap["V"] = 5
	romansMap["X"] = 10
	romansMap["L"] = 50
	romansMap["C"] = 100
	romansMap["D"] = 500
	romansMap["M"] = 1000
	return romansMap
}
