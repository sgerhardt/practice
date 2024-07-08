package search

// Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums.
// If target exists, then return its index. Otherwise, return -1.
//
// You must write an algorithm with O(log n) runtime complexity.
func search(nums []int, target int) int {
	pivot, left := 0, 0
	right := len(nums) - 1
	// keep splitting the search area in half until we find our target...
	for left <= right {
		// the pivot is the median of each iteration
		pivot = left + (right-left)/2

		if nums[pivot] == target {
			return pivot
		}

		if target < nums[pivot] {
			// shift search to the left of the slice
			right = pivot - 1
		} else {
			// shift search to the right of the slice
			left = pivot + 1
		}
	}
	return -1
}
