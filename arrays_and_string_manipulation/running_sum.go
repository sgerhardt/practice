package arrays_and_string_manipulation

// Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
//
// Return the running sum of nums.
//
// Input: nums = [1,2,3,4]
// Output: [1,3,6,10]
// Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].
func runningSum(nums []int) []int {
	results := make([]int, len(nums))

	sum := 0
	for i, num := range nums {
		sum += num
		results[i] = sum
	}
	return results
}
