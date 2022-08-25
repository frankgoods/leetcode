package problems

func fillCount(grid [][]int, i, j int) int {
	if i < 0 || i == len(grid) || j < 0 || j == len(grid[0]) || grid[i][j] == 0 {
		return 0
	}

	grid[i][j] = 0

	return 1 + fillCount(grid, i+1, j) + fillCount(grid, i, j+1) +
		fillCount(grid, i-1, j) + fillCount(grid, i, j-1)
}

func maxAreaOfIsland(grid [][]int) int {
	m := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				c := fillCount(grid, i, j)
				if c > m {
					m = c
				}
			}
		}
	}

	return m
}
