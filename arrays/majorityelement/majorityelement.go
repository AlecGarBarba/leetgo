package majorityelement

func majorityElement(nums []int) int {
	majority, frequency := 0, 0
	for _, v := range nums {

		if frequency == 0 {
			majority = v
			frequency = 1
		} else if v == majority {
			frequency++
		} else {
			frequency--
		}

	}

	return majority
}
