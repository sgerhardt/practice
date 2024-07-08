package arrays_and_string_manipulation

import "sort"

// Given a set of coin denominations i.e. 1, 5, 25, or 1,5,6,8 return the smallest number
// of coins that can be used to give a value i.e. 27 or 11.
func denominations(denoms []int, target int) []int {
	sort.Ints(denoms)
	var result []int
	for i := len(denoms); i > 0; i-- {
		largest := denoms[i-1]
		for target-largest >= 0 {
			result = append(result, largest)
			target = target - largest
		}
	}
	return result
}
