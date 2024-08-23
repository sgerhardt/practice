package arrays_and_string_manipulation

func iterCols() []int {
	grid := [][]int{{1, 2, 3}, {4, 5, 6}}

	for col := 0; col < len(grid[0]); col++ {
		for row := 0; row < len(grid); row++ {
			println(grid[row][col])
		}
	}
	return nil
}

func iterRows() []int {
	grid := [][]int{{1, 2, 3}, {4, 5, 6}}
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			println(grid[row][col])
		}
	}
	return nil
}

func iterRowsWithRanges() []int {
	grid := [][]int{{1, 2, 3}, {4, 5, 6}}
	for rowIdx, _ := range grid {
		for colIdx, _ := range grid[rowIdx] {
			println(grid[rowIdx][colIdx])
		}
	}
	return nil
}
