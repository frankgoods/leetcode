package problems

func fill(grid [][]byte, i, j int) {
	if i < 0 || i == len(grid) || j < 0 || j == len(grid[0]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'

	fill(grid, i+1, j)
	fill(grid, i, j+1)
	fill(grid, i-1, j)
	fill(grid, i, j-1)
}

func numIslands(grid [][]byte) int {
	count := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				count++
				fill(grid, i, j)
			}
		}
	}
	return count
}
