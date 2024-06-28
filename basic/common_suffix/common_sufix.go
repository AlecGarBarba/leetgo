package commonsuffix

import (
	"math"
)

/**
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".



Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.


Constraints:

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lowercase English letters.

*/

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	if len(strs) == 1 {
		return strs[0]
	}

	for _, str := range strs {
		prefix = indCommonPrefix(prefix, str)

	}
	return prefix
}

func indCommonPrefix(str1, str2 string) string {

	if str1 == str2 {
		return str1
	}

	minLen := int(math.Min(float64(len(str1)), float64(len(str2))))

	for i := 0; i < minLen; i++ {
		if str1[i] != str2[i] {
			return str1[:i]
		}
	}
	return str1[:minLen]

}
