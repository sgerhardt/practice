package arrays_and_string_manipulation

//Given a positive integer n, generate an n x n matrix filled with elements from 1 to n2 in spiral order.
//
//
//
//Example 1:
//
//
//Input: n = 3
//Output: [[1,2,3],[8,9,4],[7,6,5]]

func generateMatrix(n int) [][]int {

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	//edge case
	if n == 0 {
		return matrix
	}

	current := 1
	rowBegin := 0
	rowEnd := n - 1
	colBegin := 0
	colEnd := n - 1

	for current <= n*n {
		i := rowBegin
		j := colBegin

		// left to right
		for j = colBegin; j <= colEnd; j++ {
			matrix[rowBegin][j] = current
			current++
		}
		rowBegin++

		// top to bottom
		for i = rowBegin; i <= rowEnd; i++ {
			matrix[i][colEnd] = current
			current++
		}
		colEnd--

		// right to left
		for j = colEnd; j >= colBegin; j-- {
			matrix[rowEnd][j] = current
			current++
		}
		rowEnd--

		// bottom to top
		for i = rowEnd; i >= rowBegin; i-- {
			matrix[i][colBegin] = current
			current++
		}
		colBegin++
	}

	return matrix
}
