package arrays_and_string_manipulation

// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
//
// Note that you must do this in-place without making a copy of the array.
//
// Example 1:
//
// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]
// Example 2:
//
// Input: nums = [0]
// Output: [0]
func moveZeroes(nums []int) {
	idxCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			// move the zero to the end
			if i == len(nums)-1 {
				// zero is already at the end...
				return
			} else {
				nums = append(nums[:i], nums[i+1:]...)
			}
			nums = append(nums, 0)
			// if we moved an element, we need to back up the index by one so we don't miss a value... e.g. []int{0,1,0}
			i--
		}
		idxCount++
		if idxCount == len(nums)-1 {
			return
		}
	}
}
