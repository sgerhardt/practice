package search

// Write an efficient algorithm that searches for a value target in an m x n integer matrix matrix. This matrix has the following properties:
//
// Integers in each row are sorted in ascending from left to right.
// Integers in each column are sorted in ascending from top to bottom.
//
// Example 1:
//
// Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
// Output: true
// Example 2:
//
// Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
// Output: false
// We start search the matrix from top right corner, initialize the current position to top right corner,
// if the target is greater than the value in current position,
// then the target can not be in entire row of current position because the row is sorted,
// if the target is less than the value in current position,
// then the target can not in the entire column because the column is sorted too.
// We can rule out one row or one column each time, so the time complexity is O(m+n).
func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil {
		return false
	}

	row := 0
	col := len(matrix[0]) - 1

	for col >= 0 && row <= len(matrix)-1 {
		if target == matrix[row][col] {
			return true
		} else if target < matrix[row][col] {
			col--
		} else if target > matrix[row][col] {
			row++
		}
	}
	return false
}
