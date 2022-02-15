package main

// Given an integer array nums sorted in non-decreasing order,
// remove the duplicates in-place such that each unique element appears only once.
// The relative order of the elements should be kept the same.

// input
// 		nums []int  -- a slice of ints sorted in increasing order, contains duplcates
//
// 		1 <= nums.length <= 3 * 10^4
//		-100 <= nums[i] <= 100
// 		nums is sorted in non-decreasing order.

// output
// 		int -- k -- the amount of unique elements in nums
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	i := 1
	for {
		if i >= len(nums) {
			break
		}
		if nums[i-1] == nums[i] {
			nums = append(nums[0:i-1], nums[i:]...)
		} else {
			i++
		}
	}

	return len(nums)
}
