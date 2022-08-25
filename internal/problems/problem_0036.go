package problems

func zero(a []int) {
	for i := 0; i < len(a); i++ {
		a[i] = 0
	}
}

func isValidSudoku(board [][]byte) bool {
	digits := [10]int{}
	const size = 9

	for i := 0; i < size; i++ {
		zero(digits[:])
		for j := 0; j < size; j++ {
			if board[i][j] == '.' {
				continue
			}
			d := int(board[i][j] - '0')
			digits[d]++
			if digits[d] > 1 {
				return false
			}
		}
	}

	for j := 0; j < size; j++ {
		zero(digits[:])
		for i := 0; i < size; i++ {
			if board[i][j] == '.' {
				continue
			}
			d := int(board[i][j] - '0')
			digits[d]++
			if digits[d] > 1 {
				return false
			}
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			zero(digits[:])
			if !checkSquare(board, i, j, digits[:]) {
				return false
			}
		}
	}

	return true
}
