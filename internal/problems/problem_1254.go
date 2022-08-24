package problems

func fill(grid [][]int, i, j int) int {
	if i < 0 || i == len(grid) || j < 0 || j == len(grid[0]) {
		return 0
	}
	if grid[i][j] == 1 {
		return 1
	}

	grid[i][j] = 1
	return fill(grid, i+1, j) & fill(grid, i, j+1) & fill(grid, i-1, j) & fill(grid, i, j-1)
}

func closedIsland(grid [][]int) int {
	count := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				count += fill(grid, i, j)
			}
		}
	}

	return count
}
