package arrays_and_string_manipulation

// Write an efficient algorithm that searches for a value target in an m x n integer matrix matrix. This matrix has the following properties:
//
// Integers in each row are sorted from left to right.
// The first integer of each row is greater than the last integer of the previous row.
//
// Example 1:
//
// Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
// Output: true
func searchMatrix(matrix [][]int, target int) bool {
	// simply convert matrix to a single slice and binary search on that
	rows := len(matrix)
	cols := len(matrix[0])

	left := 0
	right := rows*cols - 1

	for left != right {
		// we will increment the left side... so make sure we update our midpoint correctly.
		mid := left + (right-left)/2
		if matrix[mid/cols][mid%cols] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	// we use division and mod to get the adjusted indexes.
	// right/cols gives us the correct row.
	// right % cols to give us the correct index into the row, e.g. given the above example, if the target = 16, right would end up as 6. 6%4 = 2, which is our index.
	return matrix[right/cols][right%cols] == target
}
