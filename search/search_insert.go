package search

// Given a sorted array of distinct integers and a target value, return the index if the target is found.
// If not, return the index where it would be if it were inserted in order.
//
// You must write an algorithm with O(log n) runtime complexity.
// Example 1:
//
// Input: nums = [1,3,5,6], target = 5
// Output: 2
// Example 2:
//
// Input: nums = [1,3,5,6], target = 2
// Output: 1
// Example 3:
//
// Input: nums = [1,3,5,6], target = 7
// Output: 4
func searchInsert(nums []int, target int) int {
	// perform a binary search...
	left := 0
	right := len(nums) - 1
	pivot := 0
	for left <= right {
		// the pivot is the median of each iteration
		pivot = left + (right-left)/2

		if target == nums[pivot] {
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
	// we haven't found our target... determine where it would have been inserted
	if left == 0 && target < nums[0] {
		// target is smaller than smallest value... return first position
		return 0
	} else if target > nums[len(nums)-1] {
		// target was not found and is greater than our values so place it at the end
		return len(nums)
	} else if len(nums) > pivot+1 && target < nums[pivot] {
		return pivot
	} else if len(nums) > pivot+1 && target < nums[pivot+1] {
		return pivot + 1
	} else if target < nums[pivot] {
		return pivot
	}
	return pivot
}
