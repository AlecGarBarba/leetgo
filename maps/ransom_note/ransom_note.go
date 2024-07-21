package ransomnote

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[string]int)

	for _, v := range magazine {
		_, exists := m[string(v)]
		if exists {
			m[string(v)]++
		} else {
			m[string(v)] = 1
		}
	}

	for _, v := range ransomNote {
		_, val := m[string(v)]
		if val {
			v2 := m[string(v)]
			if v2 == 0 {
				return false
			}
			m[string(v)]--

		} else {
			return false
		}
	}

	return true

}
