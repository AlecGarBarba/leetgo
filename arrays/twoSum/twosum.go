package twosum

/*
*
Based on the formula:
x + y = target;

It can be translated to:
y = target - x.

If we keep the value of target- x, this will give us y
Once we encounter y in our array, all we need to do is use that i + the value we store on the map
*/
func twoSum(nums []int, target int) []int {
	seenMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if value, exists := seenMap[diff]; exists {
			return []int{value, i}
		}
		seenMap[nums[i]] = i
	}
	return []int{0, 1}
}
