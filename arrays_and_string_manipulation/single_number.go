package arrays_and_string_manipulation

import "sort"

//Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
//
//You must implement a solution with a linear runtime complexity and use only constant extra space.
//Input: nums = [2,2,1]
//Output: 1

func singleNumber(nums []int) int {
	sort.Ints(nums)
	for i, num := range nums {
		// check adjacent values
		if i < len(nums)-1 && num == nums[i+1] {
			continue
		} else if i > 0 && num == nums[i-1] {
			continue
		} else {
			return num
		}
	}
	return 0
}
