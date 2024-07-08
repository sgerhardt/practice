package arrays_and_string_manipulation

import "sort"

//Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
//
//Notice that the solution set must not contain duplicate triplets.

// Input: nums = [-1,0,1,2,-1,-4]
// Output: [[-1,-1,2],[-1,0,1]]
// Explanation:
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
// The distinct triplets are [-1,0,1] and [-1,-1,2].
// Notice that the order of the output and the order of the triplets does not matter.
// Input: nums = [0,1,1]
// Output: []
// Explanation: The only possible triplet does not sum up to 0.
// Example 3:
func threeSum(nums []int) [][]int {

	sort.Ints(nums)

	var res [][]int
	//The idea is to sort an input array and then run through all indices of a possible first element of a triplet.
	//For each possible first element we make a standard bi-directional sweep of the remaining part of the array.
	//Also, we want to skip equal elements to avoid duplicates in the answer.
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			low := i + 1
			hi := len(nums) - 1
			// e.g. first number is 5, so sum is -5. The other two numbers need to add up to equal -5 so sum of all three numbers is 0.
			//E.g. the other two numbers could be -3 and -2. -3 + -2 = -5. Triplet is {-3,-2,5}
			sum := 0 - nums[i]
			for low < hi {
				if nums[low]+nums[hi] == sum {
					res = append(res, []int{nums[i], nums[low], nums[hi]})
					for low < hi && nums[low] == nums[low+1] {
						// skip duplicates
						low++
					}
					for low < hi && nums[hi] == nums[hi-1] {
						// skip duplicates
						hi--
					}
					low++
					hi--
				} else if nums[low]+nums[hi] < sum {
					low++
				} else {
					hi--
				}
			}
		}
	}

	return res
}
