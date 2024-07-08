package arrays_and_string_manipulation

import "math"

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	small := math.MaxInt
	big := math.MaxInt

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num <= small {
			// update small if num is smaller than both
			small = nums[i]
		} else if num <= big {
			// update big only if greater than small but smaller than big
			big = num
		} else {
			// return true if you find a number bigger than both
			return true
		}
	}

	return false
}
