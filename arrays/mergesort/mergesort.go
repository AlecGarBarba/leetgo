package mergesort

/*
You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	// create a copy of nums1 array
	left := make([]int, m)
	copy(left, nums1[:m])

	p1, p2 := 0, 0

	for i := 0; i < m+n; i++ {

		if p1 >= m {
			nums1[i] = nums2[p2]
			p2++
			continue

		}
		if p2 >= n || left[p1] < nums2[p2] {
			nums1[i] = left[p1]
			p1++
		} else {
			nums1[i] = nums2[p2]
			p2++
		}
	}

}
