package graphs

// The idea is to count the area of each island using dfs. During the dfs, we set the value of each point in the island to 0. The time complexity is O(mn).
func maxAreaOfIsland(grid [][]int) int {
	max_area := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				max_area = max(max_area, areaOfIsland(grid, i, j))
			}
		}
	}
	return max_area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func areaOfIsland(grid [][]int, i, j int) int {
	if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) && grid[i][j] == 1 {
		grid[i][j] = 0 // this is how we mark it as "visited". We return the int to keep track of island size
		// recursively search south, north, west and east for more land!
		return 1 + areaOfIsland(grid, i+1, j) + areaOfIsland(grid, i-1, j) + areaOfIsland(grid, i, j-1) + areaOfIsland(grid, i, j+1)
	}
	return 0
}
