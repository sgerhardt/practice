package arrays_and_string_manipulation

//Given an array nums of size n, return the majority element.
//
//The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.
//
//
//
//Example 1:
//
//Input: nums = [3,2,3]
//Output: 3
//Example 2:
//
//Input: nums = [2,2,1,1,1,2,2]
//Output: 2
//

func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	numToOccurrence := map[int]int{}
	for _, num := range nums {
		val, found := numToOccurrence[num]
		if !found {
			numToOccurrence[num] = 1
		} else {
			if val+1 > len(nums)/2 {
				return num
			}
			numToOccurrence[num] = val + 1
		}
	}
	return 0
}
