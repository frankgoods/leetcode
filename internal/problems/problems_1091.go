package problems

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}

	q := [][]int{{0, 0}, nil}
	n := len(grid)
	grid[0][0] = 10

	ans := 1
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		if cur == nil {
			if len(q) != 0 {
				q = append(q, nil)
				ans++
			}
		} else if cur[0] == n-1 && cur[1] == n-1 {
			return ans
		} else {
			i, j := cur[0], cur[1]
			children := [][]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}
			for _, ch := range children {
				di := i + ch[0]
				dj := j + ch[1]
				if dj >= 0 && dj < n && di >= 0 && di < n && grid[di][dj] == 0 {
					q = append(q, []int{di, dj})
					grid[di][dj] = 10
				}
			}
		}
	}

	return -1
}
