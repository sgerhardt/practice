package math

func findDuplicateWithMap(nums []int) int {
	numbers := map[int]struct{}{}
	for _, num := range nums {
		if _, ok := numbers[num]; ok {
			return num
		} else {
			numbers[num] = struct{}{}
		}
	}
	return 0
}

func findDuplicateWithMath(nums []int) int {
	n := len(nums) - 1

	sumWithoutDuplicate := (n*n + n) / 2

	actualSum := 0
	for _, num := range nums {
		actualSum += num
	}
	return actualSum - sumWithoutDuplicate
}
