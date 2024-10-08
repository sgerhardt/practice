package arrays_and_string_manipulation

//Given an integer array nums and an integer val, remove all occurrences of val in nums in-place.
//The relative order of the elements may be changed.
//
//Since it is impossible to change the length of the array in some languages, you must instead
//have the result be placed in the first part of the array nums.
//More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result.
//It does not matter what you leave beyond the first k elements.
//
//Return k after placing the final result in the first k slots of nums.
//
//Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.
//Input: nums = [3,2,2,3], val = 3
//Output: 2, nums = [2,2,_,_]

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if val == nums[i] && i == len(nums) {
			nums = remove(nums, i-1)
			i -= 1
		} else if val == nums[i] && i <= len(nums) {
			nums = remove(nums, i)
			i -= 1
		}
	}

	return len(nums)
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
