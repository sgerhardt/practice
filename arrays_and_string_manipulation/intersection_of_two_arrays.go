package arrays_and_string_manipulation

// Given two integer arrays nums1 and nums2, return an array of their intersection.
// Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order.
//
// Example 1:
//
// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2,2]
// Example 2:
//
// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// Output: [4,9]
// Explanation: [9,4] is also accepted.
// Runtime of N
func intersect(nums1 []int, nums2 []int) []int {
	numToOccurrences := make(map[int]int, len(nums1))

	for _, num := range nums1 {
		numToOccurrences[num]++
	}

	var intersections []int

	for _, val := range nums2 {
		occurrences, found := numToOccurrences[val]
		if found && occurrences > 0 {
			intersections = append(intersections, val)
			numToOccurrences[val]--
		}
	}
	return intersections
}
