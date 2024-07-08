package arrays_and_string_manipulation

//Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.
//
//A subarray is a contiguous non-empty sequence of elements within an array.
//
//
//
//Example 1:
//
//Input: nums = [1,1,1], k = 2
//Output: 2

// We are using a map to store the prefix sum count.
// Declare two variables one for the answer (count) and one to store the sum.
// Traverse through the array and keep on adding the elements to sum (subarray of a single element).
// Now, check if (sum-k) is present in the unordered map, if it is present it just means that a subarray exists which has a sum equal to k.
// if (sum-k) doesn't exist just insert the sum to the unordered map. as the sum is the sum of the array up to that index.
// Time complexity: O(n).
func subarraySum(nums []int, k int) int {
	sumToOccurrence := map[int]int{}

	sum := 0
	count := 0
	sumToOccurrence[sum] = 1 // e.g. target is 2. we have subarray [2]. sum - k == 0, so we should have an entry for it...

	for idx, num := range nums {
		print(idx)
		sum += num
		find := sum - k
		if val, exists := sumToOccurrence[find]; exists {
			count += val
		}
		sumToOccurrence[sum]++
	}

	return count
}
