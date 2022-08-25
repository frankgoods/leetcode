package problems

func fill(isConn [][]int, color []int, i int, col int) {
	color[i] = col
	for j, c := range isConn[i] {
		if c == 1 && color[j] != col {
			fill(isConn, color, j, col)
		}
	}
}

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)

	color := make([]int, n)

	col := 0
	for i := 0; i < n; i++ {
		if color[i] == 0 {
			col++
			fill(isConnected, color, i, col)
		}
	}
	return col
}
