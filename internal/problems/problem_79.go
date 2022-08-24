package problems

func dps(board [][]byte, i, j int, word string, strPos int) bool {
	if strPos == len(word) {
		return true
	}

	if i < 0 || j < 0 || i == len(board) ||
		j == len(board[0]) || board[i][j] != word[strPos] {
		return false
	}

	tmp := board[i][j]
	board[i][j] = '_'

	res := dps(board, i+1, j, word, strPos+1) ||
		dps(board, i-1, j, word, strPos+1) ||
		dps(board, i, j+1, word, strPos+1) ||
		dps(board, i, j-1, word, strPos+1)

	board[i][j] = tmp

	return res
}

func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[i] {
			if dps(board, i, j, word, 0) {
				return true
			}
		}
	}

	return false
}
