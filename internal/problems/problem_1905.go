package problems

func color(grid1 [][]int, i, j int, col int) {
	if i < 0 || i == len(grid1) || j < 0 || j == len(grid1[0]) ||
		grid1[i][j] != 1 {
		return
	}

	grid1[i][j] = col
	color(grid1, i+1, j, col)
	color(grid1, i, j+1, col)
	color(grid1, i-1, j, col)
	color(grid1, i, j-1, col)
}

func check(grid2, grid1 [][]int, i, j int, col int) int {
	if i < 0 || i == len(grid2) || j < 0 || j == len(grid2[0]) ||
		grid2[i][j] != 1 {
		return 1
	}

	res := 1
	if grid1[i][j] != col {
		res = 0
	}

	grid2[i][j] = 0

	return res & check(grid2, grid1, i+1, j, col) & check(grid2, grid1, i, j+1, col) & check(grid2, grid1, i-1, j, col) & check(grid2, grid1, i, j-1, col)
}

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	col := 2
	for i := range grid1 {
		for j := range grid1[i] {
			if grid1[i][j] == 1 {
				color(grid1, i, j, col)
				col++
			}
		}
	}

	count := 0
	for i := range grid2 {
		for j := range grid2[i] {
			if grid2[i][j] == 1 {
				c := grid1[i][j]
				if c == 0 {
					c = 100
				}
				count += check(grid2, grid1, i, j, c)
			}
		}
	}

	return count
}
