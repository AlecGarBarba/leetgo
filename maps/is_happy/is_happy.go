package ishappy

import "math"

func isHappy(n int) bool {
	s := make(map[int]bool)
	curr := n
	for {
		if curr == 1 {
			return true
		}
		sum := getSumOfSquaredDigits(curr)

		if ok := s[sum]; ok {
			return false
		}

		// get sum of suare of digits
		s[sum] = true
		curr = sum
	}
}

func getSumOfSquaredDigits(n int) int {
	sumOfSquaredDigits := 0
	for n > 0 {
		rightmostdigit := n % 10
		// do something with the extracted digit
		sumOfSquaredDigits += int(math.Pow(float64(rightmostdigit), 2))
		n /= 10 // discard the rightmost at go to the next
	}
	return sumOfSquaredDigits
}
