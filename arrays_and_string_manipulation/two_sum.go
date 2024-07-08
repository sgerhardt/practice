package arrays_and_string_manipulation

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
// You can return the answer in any order.
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
func twoSum(nums []int, target int) []int {
	targetDistanceToIndex := map[int]int{}
	for idx, num := range nums {
		val, ok := targetDistanceToIndex[num]
		if ok {
			return []int{val, idx}
		} else {
			targetDistance := target - num
			targetDistanceToIndex[targetDistance] = idx
		}
	}
	return []int{}
}

func twoSumSlow(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

//class Solution(object):
//def twoSum(self, nums, target):
//"""
//:type nums: List[int]
//:type target: int
//:rtype: List[int]
//"""
//targetDistanceToIndex = {}
//for idx, val in enumerate(nums):
//lookup = targetDistanceToIndex.get(val)
//if lookup is not None:
//return [idx, lookup]
//targetDistance = target - val
//targetDistanceToIndex[targetDistance] = idx
//
//return []
