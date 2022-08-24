package problems

func fill(board [][]byte, i, j int) {
	n := len(board)
	m := len(board[0])

	if i < 0 || i >= n || j < 0 || j >= m || board[i][j] != 'O' {
		return
	}
	board[i][j] = '!'
	fill(board, i, j+1)
	fill(board, i, j-1)
	fill(board, i+1, j)
	fill(board, i-1, j)
}

func solve(board [][]byte) {
	n := len(board)
	m := len(board[0])

	for i := 0; i < n; i++ {
		fill(board, i, 0)
		fill(board, i, m-1)
	}
	for j := 0; j < m; j++ {
		fill(board, 0, j)
		fill(board, n-1, j)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == '!' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}
