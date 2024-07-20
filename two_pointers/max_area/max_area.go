package maxarea

func maxArea(height []int) int {
	left, right := 0, len(height)-1

	maxArea := 0

	for left < right {

		b := right - left
		h := min(height[right], height[left])

		currArea := area(b, h)

		if currArea > maxArea {
			maxArea = currArea
		}

		if height[left] <= height[right] {
			left++
		} else {
			right--
		}

	}

	return maxArea
}

func area(b, h int) int {
	return b * h
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
