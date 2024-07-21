package isisomorphic

func isIsomorphic(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	// create a map of key: char of s, val, char of t
	mapST := make(map[byte]byte)
	mapTS := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		valST, okST := mapST[s[i]]
		valTS, okTS := mapTS[t[i]]

		// if mapping from ST exists and it's a different letter
		// word is not isomorphic
		if okST && valST != t[i] {
			return false
		}

		// if mapping from ts exists and its different from
		// s[i], word is not isomorphic
		if okTS && valTS != s[i] {
			return false
		}
		mapST[s[i]] = t[i]
		mapTS[t[i]] = s[i]
	}

	// everything has been validated
	return true

}
