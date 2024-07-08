package arrays_and_string_manipulation

//Given an integer numRows, return the first numRows of Pascal's triangle.
//
//In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:
//Input: numRows = 5
//Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	} else if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}

	output := [][]int{{1}, {1, 1}}
	for row := 2; row < numRows; row++ {
		// the length of the slice is equal to its place in outer slice + 1
		slice := make([]int, row+1)
		for i := range slice {
			// first and last entry is always 1
			if i == 0 || i == len(slice)-1 {
				slice[i] = 1
				continue
			}
			slice[i] = output[row-1][i-1] + output[row-1][i]
		}
		output = append(output, slice)
	}
	return output
}

func getRow(idx int) []int {
	if idx == 0 {
		return []int{1}
	} else if idx == 1 {
		return []int{1, 1}
	}

	triangle := [][]int{{1}, {1, 1}}
	for row := 2; row <= idx; row++ {
		// the length of the slice is equal to its place in outer slice + 1
		slice := make([]int, row+1)
		for i := range slice {
			// first and last entry is always 1
			if i == 0 || i == len(slice)-1 {
				slice[i] = 1
				continue
			}
			slice[i] = triangle[row-1][i-1] + triangle[row-1][i]
		}
		triangle = append(triangle, slice)
	}
	return triangle[idx]
}
