package problems

func fillAndCount(grid [][]int, i, j int) int {
	if i < 0 || i == len(grid) || j < 0 || j == len(grid[0]) {
		return -10000
	}

	if grid[i][j] == 0 {
		return 0
	}

	grid[i][j] = 0

	return 1 + fillAndCount(grid, i+1, j) + fillAndCount(grid, i, j+1) +
		fillAndCount(grid, i-1, j) + fillAndCount(grid, i, j-1)
}

func numEnclaves(grid [][]int) int {
	c := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				v := fillAndCount(grid, i, j)
				if v > 0 {
					c += v
				}
			}
		}
	}

	return c
}
