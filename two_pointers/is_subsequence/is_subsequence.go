package issubsequence

func isSubsequence(s string, t string) bool {
	left := 0
	for right := 0; right < len(t); right++ {

		if left == len(s) {
			return true
		}
		if s[left] == t[right] {
			left++
		}
	}
	return left >= len(s)
}
