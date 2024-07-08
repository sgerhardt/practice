package arrays_and_string_manipulation

// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
//
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
//
// You must write an algorithm that runs in O(n) time and without using the division operation.
//
// Example 1:
//
// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]
// Example 2:
//
// Input: nums = [-1,1,0,-3,3]
// Output: [0,0,9,0,0]
//
// Given numbers [2, 3, 4, 5], regarding the third number 4, the product of array except 4 is 2*3*5 which consists of two parts: left 2*3 and right 5.
// The product is left*right. We can get lefts and rights:
// Numbers:     2    3    4     5
// Lefts:            2  2*3 2*3*4
// Rights:  3*4*5  4*5    5
// Let’s fill the empty with 1:
//
// Numbers:     2    3    4     5
// Lefts:       1    2  2*3 2*3*4
// Rights:  3*4*5  4*5    5     1
// We can calculate lefts and rights in 2 loops. The time complexity is O(n).
//
// We store lefts in result array. If we allocate a new array for rights. The space complexity is O(n). To make it O(1), we just need to store it in a variable which is right.
func productExceptSelf(nums []int) []int {

	res := make([]int, len(nums))

	left := 1
	// Calculate lefts and store in res.
	for i := 0; i < len(nums); i++ {
		if i > 0 {
			left = left * nums[i-1]
		}
		res[i] = left
	}

	// Calculate the rights and the product from the end of the array
	right := 1
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		if i < n-1 {
			right = right * nums[i+1]
		}
		res[i] *= right
	}
	return res
}
