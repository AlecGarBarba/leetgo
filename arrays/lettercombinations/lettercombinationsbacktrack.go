package lettercombinations

type SliceWrapper struct {
	slice []string
}

func letterCombinationsBackTrack(digits string) []string {

	ans := &SliceWrapper{slice: []string{}}

	if len(digits) == 0 {
		return ans.slice
	}

	backtrack(0, "", digits, ans)

	return ans.slice

}

func backtrack(index int, str string, digits string, ans *SliceWrapper) {

	if index == len(digits) {
		ans.slice = append(ans.slice, str)
		return
	}
	m := getDigitsToLettersMap()

	s := m[string(digits[index])]

	for i := 0; i < len(s); i++ {
		str += s[i]
		backtrack(index+1, str, digits, ans)
		str = str[0 : len(str)-1]
	}
}
