package lettercombinations

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	m := getDigitsToLettersMap()

	if len(digits) == 1 {
		return m[string(digits[0])]
	}

	// let's do twto

	if len(digits) == 2 {
		return getDoubleDigitMapping(m, digits)

	}

	if len(digits) == 3 {
		return getTripleDigitMapping(m, digits)
	}

	return getQuadrupleDigitMapping(m, digits)
}

func getQuadrupleDigitMapping(m map[string][]string, digits string) []string {
	f := m[string(digits[0])]
	s := m[string(digits[1])]
	t := m[string(digits[2])]
	q := m[string(digits[3])]

	all := []string{}

	for _, firstv := range f {
		for _, secondv := range s {

			for _, thirdv := range t {

				for _, fourthv := range q {

					all = append(all, string(firstv)+string(secondv)+string(thirdv)+string(fourthv))
				}
			}

		}
	}

	return all
}

func getTripleDigitMapping(m map[string][]string, digits string) []string {
	f := m[string(digits[0])]
	s := m[string(digits[1])]
	t := m[string(digits[2])]

	all := []string{}

	for _, firstv := range f {
		for _, secondv := range s {

			for _, thirdv := range t {
				all = append(all, string(firstv)+string(secondv)+string(thirdv))
			}

		}
	}

	return all
}

func getDoubleDigitMapping(m map[string][]string, digits string) []string {
	f := m[string(digits[0])]
	s := m[string(digits[1])]

	all := []string{}

	for _, firstv := range f {
		for _, secondv := range s {
			all = append(all, string(firstv)+string(secondv))
		}
	}

	return all
}

func getDigitsToLettersMap() map[string][]string {
	return map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
}
