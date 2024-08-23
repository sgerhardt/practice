package arrays_and_string_manipulation

// CountingSort sorts an array of integers using counting sort algorithm.
func CountingSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	// Find the maximum element in the array
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	// Initialize the count array
	count := make([]int, max+1)

	// Count the occurrences of each element
	for _, num := range arr {
		count[num]++
	}

	// Modify the count array such that each element at each index
	// contains the sum of the previous counts
	for i := 1; i <= max; i++ {
		count[i] += count[i-1]
	}

	// Build the output array
	output := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		num := arr[i]
		output[count[num]-1] = num
		count[num]--
	}

	return output
}
